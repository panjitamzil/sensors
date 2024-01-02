package usecase

import (
	"sensors/config"
	"sensors/domain/sensors"
)

type service struct {
	sensorRepo sensors.RepoInterface
	config     *config.Config
}

type Dependencies struct {
	SensorRepo sensors.RepoInterface
	Config     *config.Config
}

func NewService(deps Dependencies) sensors.UsecaseInterface {
	return &service{
		sensorRepo: deps.SensorRepo,
		config:     deps.Config,
	}
}
