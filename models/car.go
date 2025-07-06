package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Car struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Year      string    `json:"year"`
	Brand     string    `json:"brand"`
	FUelType  string    `json:"fuel_type"`
	Engine    Engine    `json:"engine"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CarRequest struct {
	Name     string  `json:"name"`
	Year     string  `json:"year"`
	Brand    string  `json:"brand"`
	FUelType string  `json:"fuel_type"`
	Engine   Engine  `json:"engine"`
	Price    float64 `json:"price"`
}

func ValidateRequest(carRequest CarRequest) error {
	if err := validateName(carRequest.Name); err != nil {
		return err
	}
	if err := validateYear(carRequest.Year); err != nil {
		return err
	}
	if err := validateBrand(carRequest.Brand); err != nil {
		return err
	}
	if err := validateFuelType(carRequest.FUelType); err != nil {
		return err
	}
	if err := validateEngine(carRequest.Engine); err != nil {
		return err
	}
	if err := validatePrice(carRequest.Price); err != nil {
		return err
	}
	return nil
}

func validateName(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	return nil
}

func validateYear(year string) error {
	if year == "" {
		return errors.New("year cannot be empty")
	}
	_, err := strconv.Atoi(year)
	if err != nil {
		return errors.New("year must be a valid number")
	}
	currentYear := time.Now().Year()
	if yearInt, _ := strconv.Atoi(year); yearInt < 1886 || yearInt > currentYear {
		return errors.New("year must be between 1886 and the current year")
	}

	return nil
}

func validateBrand(brand string) error {
	if brand == "" {
		return errors.New("brand cannot be empty")
	}
	return nil
}

func validateFuelType(fuelType string) error {
	if fuelType == "" {
		return errors.New("fuel type cannot be empty")
	}
	validFuelTypes := []string{"Petrol", "Diesel", "Electric", "Hybrid"}
	for _, validType := range validFuelTypes {
		if fuelType == validType {
			return nil
		}
	}
	return errors.New("fuel type must be one of: Petrol, Diesel, Electric, Hybrid")
}

func validateEngine(engine Engine) error {
	if engine.EngineID == uuid.Nil {
		return errors.New("engine ID cannot be empty")
	}
	if engine.Displacement <= 0 {
		return errors.New("engine displacement must be greater than zero")
	}
	if engine.NoOfCylinders <= 0 {
		return errors.New("number of cylinders must be greater than zero")
	}
	if engine.CarRange <= 0 {
		return errors.New("car range must be greater than zero")
	}
	return nil
}

func validatePrice(price float64) error {
	if price <= 0 {
		return errors.New("price must be greater than zero")
	}
	return nil
}
