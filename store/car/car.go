package car

import (
	"context"
	"database/sql"

	"github.com/diegobbrito/car-zone/models"
)

type Store struct {
	db *sql.DB
}

func new(db *sql.DB) Store {
	return Store{db: db}
}

func (s Store) GetCarById(ctx context.Context, id string) (models.Car, error) {
	var car models.Car
	query := `SELECT c.id, c.name, c.year, c.brand, c.fuel_type, c.engine_id, c.price, c.created_at, c.updated_at,
	 	e.id, e.displacement, e.no_of_cylinders, e.car_range 
		FROM car c LEFT JOIN engine e ON c.engine_id = e.id
		WHERE c.id = $1`
	row := s.db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&car.ID, &car.Name, &car.Year, &car.Brand, &car.FuelType, &car.Engine.EngineID, &car.Price, &car.CreatedAt, &car.UpdatedAt,
		&car.Engine.EngineID, &car.Engine.Displacement, &car.Engine.NoOfCylinders, &car.Engine.CarRange,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return car, nil
		}
		return car, err
	}
	return car, nil
}

func (s Store) GetCarByBrand(ctx context.Context, brand string, isEngine bool) ([]models.Car, error) {
	var cars []models.Car
	var query string
	if isEngine {
		query = `SELECT c.id, c.name, c.year, c.brand, c.fuel_type, c.engine_id, c.price, c.created_at, c.updated_at,
			e.id, e.displacement, e.no_of_cylinders, e.car_range 
			FROM car c LEFT JOIN engine e ON c.engine_id = e.id
			WHERE c.brand = $1`
	} else {
		query = `SELECT id, name, year, brand, fuel_type, engine_id, price, created_at, updated_at FROM car WHERE brand = $1`
	}
	rows, err := s.db.QueryContext(ctx, query, brand)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var car models.Car
		if isEngine {
			var engine models.Engine
			err = rows.Scan(
				&car.ID, &car.Name, &car.Year, &car.Brand, &car.FuelType, &car.Engine.EngineID, &car.Price, &car.CreatedAt, &car.UpdatedAt,
				&car.Engine.EngineID, &car.Engine.Displacement, &car.Engine.NoOfCylinders, &car.Engine.CarRange,
			)
			if err != nil {
				return nil, err
			}
			car.Engine = engine
		} else {
			err = rows.Scan(&car.ID, &car.Name, &car.Year, &car.Brand, &car.FuelType, &car.Engine.EngineID, &car.Price, &car.CreatedAt, &car.UpdatedAt)
			if err != nil {
				return nil, err
			}
		}

		cars = append(cars, car)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return cars, nil
}

func (s Store) CreateCar(cxt context.Context, carRequest *models.CarRequest) (models.Car, error) {

}

func (s Store) UpdateeCar(cxt context.Context, id string, carRequest *models.CarRequest) (models.Car, error) {

}

func (s Store) DeleteCar(cxt context.Context, id string) (models.Car, error) {

}
