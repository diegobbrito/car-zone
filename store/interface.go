package store

import (
	"context"

	"github.com/diegobbrito/car-zone/models"
)

type CarStoreInterface interface {
	GetCarById(ctx context.Context, id string) (models.Car, error)
	GetCarByBrand(ctx context.Context, brand string, isEngine bool) ([]models.Car, error)
	CreateCar(ctx context.Context, carRequest *models.CarRequest) (models.Car, error)
	UpdateCar(ctx context.Context, id string, carRequest *models.CarRequest) (models.Car, error)
	DeleteCar(ctx context.Context, id string) (models.Car, error)
}

type EngineStoreInterface interface {
	GetEngineById(ctx context.Context, id string) (models.Engine, error)
	CreateEngine(ctx context.Context, engineRequest *models.EngineRequest) (models.Engine, error)
	UpdateEngine(ctx context.Context, id string, engineRequest *models.EngineRequest) (models.Engine, error)
	DeleteEngine(ctx context.Context, id string) (models.Engine, error)
}
