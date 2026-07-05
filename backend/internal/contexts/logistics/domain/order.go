package domain

import "time"

type Order struct {
	ID              string    `json:"id"`
	CustomerID      string    `json:"customer_id"`
	PickupAddress   string    `json:"pickup_address"`
	DeliveryAddress string    `json:"delivery_address"`
	Status          string    `json:"status"` // PENDING, ASSIGNED, IN_TRANSIT, DELIVERED
	CreatedAt       time.Time `json:"created_at"`
}