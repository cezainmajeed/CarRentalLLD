package models

type VehicleType string

const (
	Sedan     VehicleType = "SEDAN"
	SUV       VehicleType = "SUV"
	Hatchback VehicleType = "HATCHBACK"
)

type VehicleStatus string

const (
	VehicleAvailable VehicleStatus = "AVAILABLE"
	VehicleRented    VehicleStatus = "RENTED"
)

type Vehicle struct {
	Id                 string
	RegistrationNumber string
	Type               VehicleType
	Status             VehicleStatus
	RatePerHour        float64
}

func (v *Vehicle) IsAvailable() bool {
	return v.Status == VehicleAvailable
}

func (v *Vehicle) MarkRented() {
	v.Status = VehicleRented
}

func (v *Vehicle) MarkAvailable() {
	v.Status = VehicleAvailable
}
