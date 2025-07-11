package service

import (
	"context"

	"github.com/diegobbrito/car-zone/models"
)

type CarServiceInterface interface {
	GetCarById(ctx context.Context, id string) (models.Car, error)
	GetCarByBrand(ctx context.Context, brand string, isEngine bool) ([]models.Car, error)
	CreateCar(ctx context.Context, carRequest *models.CarRequest) (models.Car, error)
	UpdateCar(ctx context.Context, id string, carRequest *models.CarRequest) (models.Car, error)
	DeleteCar(ctx context.Context, id string) (models.Car, error)
}
