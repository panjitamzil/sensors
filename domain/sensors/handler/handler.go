package handler

import (
	"context"
	sensor "sensors/domain/sensors"
	sensorpb "sensors/pb"
)

type Handler struct {
	usecase sensor.UsecaseInterface
}

func NewHandler(usecase sensor.UsecaseInterface) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) InsertSensor(ctx context.Context, in *sensorpb.PayloadRequest) (*sensorpb.MessageReturn, error) {
	err := h.usecase.InsertSensor(in)
	if err != nil {
		return &sensorpb.MessageReturn{
			Message: "Failed to insert sensor",
		}, err
	}

	return &sensorpb.MessageReturn{
		Message: "Success",
	}, nil
}

func (h *Handler) GetSensors(ctx context.Context, in *sensorpb.PayloadGetSensors) (*sensorpb.Sensors, error) {
	var datas []*sensorpb.Sensor

	result, err := h.usecase.GetSensors(in)
	if err != nil {
		return nil, err
	}

	for _, value := range result {
		datas = append(datas, value.ToPB())
	}

	var list = sensorpb.Sensors{
		Sensors: datas,
	}

	return &list, nil
}

func (h *Handler) EditSensor(ctx context.Context, in *sensorpb.PayloadUpdateSensors) (*sensorpb.MessageReturn, error) {
	err := h.usecase.EditSensors(in)
	if err != nil {
		return &sensorpb.MessageReturn{
			Message: "Failed to edit sensor",
		}, err
	}

	return &sensorpb.MessageReturn{
		Message: "Success",
	}, nil
}

func (h *Handler) DeleteSensor(ctx context.Context, in *sensorpb.PayloadGetSensors) (*sensorpb.MessageReturn, error) {
	err := h.usecase.DeleteSensors(in)
	if err != nil {
		return &sensorpb.MessageReturn{
			Message: "Failed to delete sensor",
		}, err
	}

	return &sensorpb.MessageReturn{
		Message: "Success",
	}, nil
}
