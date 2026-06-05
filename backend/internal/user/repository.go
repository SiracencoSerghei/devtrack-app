package user

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/db"
)

type UserStore struct {
	storage *db.PostgresDB
}

func NewUserStore(storage *db.PostgresDB) *UserStore {
	return &UserStore{storage: storage}
}

func (s *UserStore) CreateUser(name, email, passwordHash string) error {
	query := `INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3)`
	
	_, err := s.storage.DB.Exec(query, name, email, passwordHash)
	if err != nil {
		return fmt.Errorf("fallimento inserimento utente in Postgres: %w", err)
	}
	
	return nil
}

func (s *UserStore) GetUserByEmail(email string) (*db.UserField, error) {
	query := `SELECT id, name, email, password_hash FROM users WHERE email = $1`
	
	var u db.UserField
	err := s.storage.DB.QueryRow(query, email).Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash)
	
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil 
		}
		return nil, fmt.Errorf("errore durante la query di selezione utente: %w", err)
	}
	
	return &u, nil
}