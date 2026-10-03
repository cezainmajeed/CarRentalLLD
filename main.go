package main

import (
	"car-rental-system/models"
	"car-rental-system/repository"
	"car-rental-system/service"
	"car-rental-system/strategy"
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("Car Rental System")

	vehicleRepo := repository.NewInMemoryVehicleRepository()
	rentalRepo := repository.NewInMemoryRentalRepository()
	customerRepo := repository.NewInMemoryCustomerRepository()

	pricingStrategy := strategy.NewHourlyPricingStrategy()

	rentalService := service.NewRentalService(vehicleRepo, rentalRepo, customerRepo, pricingStrategy)

	//create customer
	customer := &models.Customer{
		Id:   "customer-1",
		Name: "Cezain",
	}

	if err := customerRepo.Save(customer); err != nil {
		log.Fatal(err)
	}

	// create vehicles

	vehicle := &models.Vehicle{
		Id:                 "vehicle-1",
		RegistrationNumber: "abc123",
		Type:               models.Sedan,
		Status:             models.VehicleAvailable,
		RatePerHour:        200,
	}

	if err := vehicleRepo.Save(vehicle); err != nil {
		log.Fatal(err)
	}

	// rent

	start := time.Now()
	end := time.Now().Add(10 * time.Hour)

	rental, err := rentalService.RentVehicle(customer.Id, vehicle.Id, start, end)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Rental Created :: Id : %s, Amount : %f", rental.Id, rental.TotalAmount)

	// return

	if err := rentalService.ReturnVehicle(rental.Id); err != nil {
		log.Fatal(err)
	}

	log.Printf("Vehicle returned")
}
