package driver

import "time"

type Driver struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	LicenseNumber string    `json:"license_number"`
	Phone         string    `json:"phone"`
	Status        string    `json:"status"` // AVAILABLE, IN_TRANSIT, OFF_DUTY
	UpdatedAt     time.Time `json:"updated_at"`
}