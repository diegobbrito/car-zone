package models

import (
	"github.com/google/uuid"
)

type Engine struct {
	EngineID      uuid.UUID `json:"engine_id"`
	Displacement  int64     `json:"displacement"`
	NoOfCylinders int       `json:"no_of_cylinders"`
	CarRange      int64     `json:"car_range"`
}
