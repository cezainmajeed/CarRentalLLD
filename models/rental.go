package models

import "time"

type RentalStatus string

const (
	RentalStatusActive    RentalStatus = "ACTIVE"
	RentalStatusCompleted RentalStatus = "COMPLETED"
)

type Rental struct {
	Id          string
	CustomerId  string
	VehicleId   string
	StartTime   time.Time
	EndTime     time.Time
	Status      RentalStatus
	TotalAmount float64
}

func (r *Rental) MarkComplete() {
	r.Status = RentalStatusCompleted
}
