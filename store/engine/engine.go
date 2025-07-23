package engine

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/diegobbrito/car-zone/models"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
)

type EngineStore struct {
	db *sql.DB
}

func New(db *sql.DB) *EngineStore {
	return &EngineStore{db: db}
}

func (e EngineStore) GetEngineById(ctx context.Context, id string) (models.Engine, error) {
	tracer := otel.Tracer("EngineStore")
	ctx, span := tracer.Start(ctx, "GetEngineById-Store")
	defer span.End()
	var engine models.Engine
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return engine, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Printf("Transaction rollback error: %v", rbErr)
			}
		} else {
			if cmErr := tx.Commit(); cmErr != nil {
				fmt.Printf("Transaction commit error: %v", cmErr)
			}
		}
	}()
	err = tx.QueryRowContext(ctx, "SELECT id, displacement, no_of_cylinders, car_range FROM engine WHERE id = $1", id).Scan(
		&engine.EngineID, &engine.Displacement, &engine.NoOfCylinders, &engine.CarRange,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return engine, nil
		}
		return engine, err
	}
	return engine, nil
}

func (e EngineStore) CreateEngine(ctx context.Context, engineRequest *models.EngineRequest) (models.Engine, error) {
	tracer := otel.Tracer("EngineStore")
	ctx, span := tracer.Start(ctx, "CreateEngine-Store")
	defer span.End()
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return models.Engine{}, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Printf("Transaction rollback error: %v", rbErr)
			}
		} else {
			if cmErr := tx.Commit(); cmErr != nil {
				fmt.Printf("Transaction commit error: %v", cmErr)
			}
		}
	}()

	engineID := uuid.New()
	_, err = tx.ExecContext(ctx, "INSERT INTO engine (id, displacement, no_of_cylinders, car_range) VALUES ($1, $2, $3, $4)",
		engineID, engineRequest.Displacement, engineRequest.NoOfCylinders, engineRequest.CarRange)
	if err != nil {
		return models.Engine{}, err
	}
	engine := models.Engine{
		EngineID:      engineID,
		Displacement:  engineRequest.Displacement,
		NoOfCylinders: engineRequest.NoOfCylinders,
		CarRange:      engineRequest.CarRange,
	}
	return engine, nil
}

func (e EngineStore) UpdateEngine(ctx context.Context, id string, engineRequest *models.EngineRequest) (models.Engine, error) {
	tracer := otel.Tracer("EngineStore")
	ctx, span := tracer.Start(ctx, "UpdateEngine-Store")
	defer span.End()
	engineID, err := uuid.Parse(id)
	if err != nil {
		return models.Engine{}, fmt.Errorf("invalid engine ID: %w", err)
	}
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return models.Engine{}, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Printf("Transaction rollback error: %v", rbErr)
			}
		} else {
			if cmErr := tx.Commit(); cmErr != nil {
				fmt.Printf("Transaction commit error: %v", cmErr)
			}
		}
	}()
	results, err := tx.ExecContext(ctx, "UPDATE engine SET displacement = $1, no_of_cylinders = $2, car_range = $3 WHERE id = $4",
		engineRequest.Displacement, engineRequest.NoOfCylinders, engineRequest.CarRange, engineID)

	if err != nil {
		return models.Engine{}, err
	}
	rowsAffected, err := results.RowsAffected()
	if err != nil {
		return models.Engine{}, err
	}
	if rowsAffected == 0 {
		return models.Engine{}, errors.New("No Rows Were Updated")
	}
	engine := models.Engine{
		EngineID:      engineID,
		Displacement:  engineRequest.Displacement,
		NoOfCylinders: engineRequest.NoOfCylinders,
		CarRange:      engineRequest.CarRange,
	}
	return engine, nil
}

func (e EngineStore) DeleteEngine(ctx context.Context, id string) (models.Engine, error) {
	tracer := otel.Tracer("EngineStore")
	ctx, span := tracer.Start(ctx, "DeleteEngine-Store")
	defer span.End()
	var engine models.Engine
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return models.Engine{}, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Printf("Transaction rollback error: %v", rbErr)
			}
		} else {
			if cmErr := tx.Commit(); cmErr != nil {
				fmt.Printf("Transaction commit error: %v", cmErr)
			}
		}
	}()

	err = tx.QueryRowContext(ctx, "SELECT id, displacement, no_of_cylinders, car_range FROM engine WHERE id = $1", id).Scan(
		&engine.EngineID, &engine.Displacement, &engine.NoOfCylinders, &engine.CarRange,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return engine, nil
		}
		return engine, err
	}

	result, err := tx.ExecContext(ctx, "DELETE FROM engine WHERE id = $1", id)
	if err != nil {
		return models.Engine{}, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Engine{}, err
	}
	if rowsAffected == 0 {
		return models.Engine{}, errors.New("No Rows Were Deleted")
	}
	return engine, nil
}
