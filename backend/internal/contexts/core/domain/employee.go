package domain

import "time"

type Company struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Employee struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	DepartmentID  string    `json:"department_id,omitempty"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Phone         string    `json:"phone"`
	RoleInCompany string    `json:"role_in_company"`
	CreatedAt     time.Time `json:"created_at"`
}