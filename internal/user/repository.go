package user

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"sync"
)

// UserRepository interface for testable service
type UserRepository interface {
	Create(u User) (User, error)
	GetAll() []User
}

// InMemory repo only for tests/dev
type Repository struct {
	mu     sync.RWMutex
	users  map[string]User
	nextID int
}

func NewRepository() *Repository {
	return &Repository{
		users:  make(map[string]User),
		nextID: 1,
	}
}

func (r *Repository) Create(u User) (User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, existing := range r.users {
		if existing.Email == u.Email {
			return User{}, errors.New("email already exists")
		}
	}

	u.ID = fmt.Sprintf("%d", r.nextID)
	r.nextID++
	r.users[u.ID] = u

	return u, nil
}

func (r *Repository) GetAll() []User {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]User, 0, len(r.users))
	for _, u := range r.users {
		result = append(result, u)
	}

	sort.Slice(result, func(i, j int) bool {
		left, _ := strconv.Atoi(result[i].ID)
		right, _ := strconv.Atoi(result[j].ID)
		return left < right
	})
	return result
}