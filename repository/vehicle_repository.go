package repository

import (
	"car-rental-system/models"
	"errors"
	"sync"
)

type VehicleRepository interface {
	GetById(id string) (*models.Vehicle, error)
	GetAvailable(vehicleType models.VehicleType) ([]*models.Vehicle, error)
	Reserve(id string) (*models.Vehicle, error)
	Free(id string) (*models.Vehicle, error)
	Save(vehicle *models.Vehicle) error
}

type InMemoryVehicleRepository struct {
	mu       sync.RWMutex
	vehicles map[string]*models.Vehicle
}

func NewInMemoryVehicleRepository() *InMemoryVehicleRepository {
	return &InMemoryVehicleRepository{
		vehicles: make(map[string]*models.Vehicle),
	}
}

func (r *InMemoryVehicleRepository) GetById(id string) (*models.Vehicle, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	vehicle, found := r.vehicles[id]
	if !found {
		return nil, errors.New("no vehicle found")
	}

	return vehicle, nil
}

func (r *InMemoryVehicleRepository) GetAvailable(vehicleType models.VehicleType) ([]*models.Vehicle, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*models.Vehicle

	for _, vehicle := range r.vehicles {
		if vehicle.Type == vehicleType && vehicle.Status == models.VehicleAvailable {
			result = append(result, vehicle)
		}
	}

	return result, nil
}

func (r *InMemoryVehicleRepository) Reserve(id string) (*models.Vehicle, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	vehicle, found := r.vehicles[id]
	if !found {
		return nil, errors.New("no vehicle found")
	}

	if !vehicle.IsAvailable() {
		return nil, errors.New("vehicle not available")
	}

	vehicle.MarkRented()

	return vehicle, nil
}

func (r *InMemoryVehicleRepository) Free(id string) (*models.Vehicle, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	vehicle, found := r.vehicles[id]
	if !found {
		return nil, errors.New("no vehicle found")
	}

	if vehicle.IsAvailable() {
		return nil, errors.New("vehicle already available")
	}

	vehicle.MarkAvailable()

	return vehicle, nil
}

func (r *InMemoryVehicleRepository) Save(vehicle *models.Vehicle) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, found := r.vehicles[vehicle.Id]
	if found {
		return errors.New("vehicle already exists")
	}

	r.vehicles[vehicle.Id] = vehicle
	return nil
}
