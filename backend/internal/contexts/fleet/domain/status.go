package domain

type DriverStatus string

const (
	StatusAvailable DriverStatus = "AVAILABLE"
	StatusInTransit DriverStatus = "IN_TRANSIT"
	StatusOffDuty   DriverStatus = "OFF_DUTY"
)