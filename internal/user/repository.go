package user

import (
	"fmt"
	"sync"
	"errors"
)

type Repository struct {
	mu    sync.RWMutex
	users map[string]User
}

func NewRepository() *Repository {
	return &Repository{
		users: make(map[string]User),
	}
}

func (r *Repository) Save(user User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[user.ID] = user
	return nil
}

func (r *Repository) GetAll() []User {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]User, 0, len(r.users))

	for _, user := range r.users {
		result = append(result, user)
	}

	return result
}

func (r *Repository) ExistsByEmail(email string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, u := range r.users {
		if u.Email == email {
			return true
		}
	}

	return false
}

func (r *Repository) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.users)
}

func (r *Repository) NextID() string {
	return fmt.Sprintf("%d", r.Count()+1)
}

func (r *Repository) Create(u User) (User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// перевірка на дублікат email
	for _, user := range r.users {
		if user.Email == u.Email {
			return User{}, errors.New("email already exists")
		}
	}

	// призначаємо ID через NextID()
	u.ID = r.NextID()

	// зберігаємо
	r.users[u.ID] = u
	return u, nil
}