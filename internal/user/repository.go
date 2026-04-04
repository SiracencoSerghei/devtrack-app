package user

import (
    "errors"
    "sort"
    "sync"

    "github.com/google/uuid"
)

type Repository interface {
    Create(u User) (User, error)
    GetAll() []User
}

type InMemoryRepository struct {
    mu    sync.RWMutex
    users map[string]User
}

func NewInMemoryRepository() *InMemoryRepository {
    return &InMemoryRepository{
        users: make(map[string]User),
    }
}

func (r *InMemoryRepository) Create(u User) (User, error) {
    r.mu.Lock()
    defer r.mu.Unlock()

    for _, existing := range r.users {
        if existing.Email == u.Email {
            return User{}, errors.New("email already exists")
        }
    }

    u.ID = uuid.NewString()
    r.users[u.ID] = u

    return u, nil
}

func (r *InMemoryRepository) GetAll() []User {
    r.mu.RLock()
    defer r.mu.RUnlock()

    result := make([]User, 0, len(r.users))
    for _, u := range r.users {
        result = append(result, u)
    }

    sort.Slice(result, func(i, j int) bool {
        return result[i].Name < result[j].Name
    })

    return result
}