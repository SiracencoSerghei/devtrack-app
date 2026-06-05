package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)


type UserRepository interface {
	CreateUser(name, email, passwordHash string) error
	GetUserByEmail(email string) (*UserField, error)
}


type UserField struct {
	ID           int
	Name         string
	Email        string
	PasswordHash string
}


type PostgresDB struct {
	DB *sql.DB
}

func NewPostgresDB(dataSourceName string) *PostgresDB {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("Impossibile aprire la connessione al database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Database non raggiungibile (Ping fallito): %v", err)
	}

	log.Println("Connessione a PostgreSQL stabilita con successo!")
	return &PostgresDB{DB: db}
}