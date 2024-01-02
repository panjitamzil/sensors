package sensors

import "sensors/domain/sensors/model"

type RepoInterface interface {
	Insert(payload model.SensorData) error
	Update(payload model.SensorData, filter model.FilterData) error
	Delete(filter model.FilterData) error
	GetAll(params *model.FilterData) (sensors []model.SensorData, err error)
}
