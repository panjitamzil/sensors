package sensors

import (
	"sensors/domain/sensors/model"
	sensorpb "sensors/pb"
)

type UsecaseInterface interface {
	InsertSensor(in *sensorpb.PayloadRequest) error
	GetSensors(in *sensorpb.PayloadGetSensors) (result []model.SensorData, err error)
	EditSensors(in *sensorpb.PayloadUpdateSensors) (err error)
	DeleteSensors(in *sensorpb.PayloadGetSensors) (err error)
}
