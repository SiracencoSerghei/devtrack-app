package main

import (
	"encoding/json"
	"log"
	"net/http"
	"golang.org/x/crypto/bcrypt"

	"github.com/SiracencoSerghei/devtrack-app/backend/config"
    "github.com/SiracencoSerghei/devtrack-app/backend/internal/db"
    "github.com/SiracencoSerghei/devtrack-app/backend/internal/user"
)

type Application struct {
	config *config.Config
	users  db.UserRepository
}

type JSONErrorResponse struct {
	Error string `json:"error"`
}

func main() {

	cfg := config.LoadConfig()

	postgresStorage := db.NewPostgresDB(cfg.DatabaseURL)
	defer postgresStorage.DB.Close()

	userStore := user.NewUserStore(postgresStorage)

	app := &Application{
		config: cfg,
		users:  userStore,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/api/signup", app.handleSignUp)
	mux.HandleFunc("/api/login", app.handleLogin)

	corsHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		mux.ServeHTTP(w, r)
	})

	log.Printf("Server DEVTRACK v2 avviato sulla porta %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, corsHandler))
}


func (app *Application) handleSignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		app.writeJSONError(w, "Metodo non consentito", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		app.writeJSONError(w, "Payload JSON non valido", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Errore Bcrypt: %v", err)
		app.writeJSONError(w, "Errore interno del server", http.StatusInternalServerError)
		return
	}

	err = app.users.CreateUser(input.Name, input.Email, string(hashedPassword))
	if err != nil {
		log.Printf("Errore DB in Signup: %v", err)
		
		app.writeJSONError(w, "L'indirizzo email è già registrato", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Utente registrato con successo"})
}

func (app *Application) handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		app.writeJSONError(w, "Metodo non consentito", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		app.writeJSONError(w, "Payload non valido", http.StatusBadRequest)
		return
	}

	dbUser, err := app.users.GetUserByEmail(input.Email)
	if err != nil {
		log.Printf("Errore DB in Login: %v", err)
		app.writeJSONError(w, "Errore interno del server", http.StatusInternalServerError)
		return
	}

	if dbUser == nil {
		app.writeJSONError(w, "Credenziali non valide", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.PasswordHash), []byte(input.Password))
	if err != nil {
		
		app.writeJSONError(w, "Credenziali non valide", http.StatusUnauthorized)
		return
	}

	response := map[string]interface{}{
		"access_token": "finto_jwt_token_fase_2",
		"user": map[string]string{
			"name":  dbUser.Name,
			"email": dbUser.Email,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (app *Application) writeJSONError(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(JSONErrorResponse{Error: msg})
}