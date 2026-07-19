package domain

import "time"

type Driver struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	LicenseNumber string    `json:"license_number"`
	Phone         string    `json:"phone"`
	Status        DriverStatus `json:"status"`
	UpdatedAt     time.Time `json:"updated_at"`
}