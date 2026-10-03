package repository

import (
	"car-rental-system/models"
	"errors"
	"sync"
)

type CustomerRepository interface {
	Save(customer *models.Customer) error
}

type InMemoryCustomerRepository struct {
	mu        sync.RWMutex
	customers map[string]*models.Customer
}

func NewInMemoryCustomerRepository() *InMemoryCustomerRepository {
	return &InMemoryCustomerRepository{
		customers: make(map[string]*models.Customer),
	}
}

func (r *InMemoryCustomerRepository) Save(customer *models.Customer) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, found := r.customers[customer.Id]
	if found {
		return errors.New("customer already exists")
	}

	r.customers[customer.Id] = customer

	return nil
}
