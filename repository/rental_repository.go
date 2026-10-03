package repository

import (
	"car-rental-system/models"
	"errors"
	"sync"
)

type RentalRepository interface {
	Save(rental *models.Rental) error
	GetById(id string) (*models.Rental, error)
	MarkComplete(id string) (*models.Rental, error)
}

type InMemoryRentalRepository struct {
	mu      sync.RWMutex
	rentals map[string]*models.Rental
}

func NewInMemoryRentalRepository() *InMemoryRentalRepository {
	return &InMemoryRentalRepository{
		rentals: make(map[string]*models.Rental),
	}
}

func (r *InMemoryRentalRepository) Save(rental *models.Rental) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, found := r.rentals[rental.Id]
	if found {
		return errors.New("rental already exists")
	}

	r.rentals[rental.Id] = rental
	return nil
}

func (r *InMemoryRentalRepository) GetById(id string) (*models.Rental, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	rental, found := r.rentals[id]
	if !found {
		return nil, errors.New("no rental found")
	}

	return rental, nil
}

func (r *InMemoryRentalRepository) MarkComplete(id string) (*models.Rental, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	rental, found := r.rentals[id]
	if !found {
		return nil, errors.New("no rental found")
	}

	rental.MarkComplete()

	return rental, nil
}
