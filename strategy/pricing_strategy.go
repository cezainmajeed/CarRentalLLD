package strategy

import "car-rental-system/models"

type PricingStrategy interface {
	Calculate(rental *models.Rental, vehicle *models.Vehicle) float64
}

type HourlyPricingStrategy struct{}

func NewHourlyPricingStrategy() *HourlyPricingStrategy {
	return &HourlyPricingStrategy{}
}

func (p *HourlyPricingStrategy) Calculate(rental *models.Rental, vehicle *models.Vehicle) float64 {
	duration := rental.EndTime.Sub(rental.StartTime)

	hours := duration.Hours()

	return hours * vehicle.RatePerHour
}
