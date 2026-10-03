package service

import (
	"car-rental-system/models"
	"car-rental-system/repository"
	"car-rental-system/strategy"
	"errors"
	"fmt"
	"time"
)

type RentalServiceInterface interface {
	GetAvailableVehicles(vehicleType models.VehicleType) ([]*models.Vehicle, error)
	RentVehicle(customerId string, vehicleId string, startTime time.Time, endTime time.Time) (*models.Rental, error)
	ReturnVehicle(rentalId string) error
}

type RentalService struct {
	vehicleRepo     repository.VehicleRepository
	rentalRepo      repository.RentalRepository
	customerRepo    repository.CustomerRepository
	pricingStrategy strategy.PricingStrategy
}

func NewRentalService(
	vehicleRepo repository.VehicleRepository,
	rentalRepo repository.RentalRepository,
	customerRepo repository.CustomerRepository,
	pricingStrategy strategy.PricingStrategy,
) *RentalService {
	return &RentalService{
		vehicleRepo:     vehicleRepo,
		rentalRepo:      rentalRepo,
		customerRepo:    customerRepo,
		pricingStrategy: pricingStrategy,
	}
}

func (s *RentalService) GetAvailableVehicles(vehicleType models.VehicleType) ([]*models.Vehicle, error) {
	return s.vehicleRepo.GetAvailable(vehicleType)
}

func (s *RentalService) RentVehicle(customerId string, vehicleId string, startTime time.Time, endTime time.Time) (*models.Rental, error) {
	vehicle, err := s.vehicleRepo.Reserve(vehicleId)
	if err != nil {
		return nil, err
	}

	rental := &models.Rental{
		Id:         fmt.Sprintf("rental-%d", time.Now().UnixNano()),
		CustomerId: customerId,
		VehicleId:  vehicleId,
		StartTime:  startTime,
		EndTime:    endTime,
		Status:     models.RentalStatusActive,
	}

	rental.TotalAmount = s.pricingStrategy.Calculate(rental, vehicle)

	err = s.rentalRepo.Save(rental)
	if err != nil {
		s.vehicleRepo.Free(vehicleId)
		return nil, errors.New("error while renting vehicle")
	}

	return rental, nil
}

func (s *RentalService) ReturnVehicle(rentalId string) error {
	rental, err := s.rentalRepo.GetById(rentalId)
	if err != nil {
		return err
	}

	if rental.Status == models.RentalStatusCompleted {
		return errors.New("vehicle already returned")
	}

	// mark status as complete
	s.rentalRepo.MarkComplete(rentalId)
	s.vehicleRepo.Free(rental.VehicleId)

	return nil

}
