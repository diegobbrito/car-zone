package engine

import (
	"context"

	"github.com/diegobbrito/car-zone/models"
	"github.com/diegobbrito/car-zone/store"
	"go.opentelemetry.io/otel"
)

type EngineService struct {
	store store.EngineStoreInterface
}

func NewEngineService(store store.EngineStoreInterface) *EngineService {
	return &EngineService{store: store}
}

func (s *EngineService) GetEngineByID(ctx context.Context, id string) (*models.Engine, error) {
	tracer := otel.Tracer("EngineService")
	ctx, span := tracer.Start(ctx, "GetEngineByID-Service")
	defer span.End()
	engine, err := s.store.GetEngineById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &engine, nil
}

func (s *EngineService) CreateEngine(ctx context.Context, engineRequest *models.EngineRequest) (*models.Engine, error) {
	tracer := otel.Tracer("EngineService")
	ctx, span := tracer.Start(ctx, "CreateEngine-Service")
	defer span.End()
	if err := models.ValidateEngineRequest(*engineRequest); err != nil {
		return nil, err
	}
	createdEngine, err := s.store.CreateEngine(ctx, engineRequest)
	if err != nil {
		return nil, err
	}
	return &createdEngine, nil
}

func (s *EngineService) UpdateEngine(ctx context.Context, id string, engineRequest *models.EngineRequest) (*models.Engine, error) {
	tracer := otel.Tracer("EngineService")
	ctx, span := tracer.Start(ctx, "UpdateEngine-Service")
	defer span.End()
	if err := models.ValidateEngineRequest(*engineRequest); err != nil {
		return nil, err
	}
	updatedEngine, err := s.store.UpdateEngine(ctx, id, engineRequest)
	if err != nil {
		return nil, err
	}
	return &updatedEngine, nil
}

func (s *EngineService) DeleteEngine(ctx context.Context, id string) (*models.Engine, error) {
	tracer := otel.Tracer("EngineService")
	ctx, span := tracer.Start(ctx, "DeleteEngine-Service")
	defer span.End()
	deletedEngine, err := s.store.DeleteEngine(ctx, id)
	if err != nil {
		return nil, err
	}
	return &deletedEngine, nil
}
