package handler

import (
	sensorpb "sensors/pb"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	sensorGRPC sensorpb.SensorServicesClient
}

type HandlerInterface interface {
	Insert(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	Get(c echo.Context) error
}

func NewHandler(sensorGRPC sensorpb.SensorServicesClient) *Handler {
	return &Handler{
		sensorGRPC: sensorGRPC,
	}
}
