package handler

import (
	"context"
	"log"
	"net/http"
	"sensors/domain/sensors-api/model"
	sensorpb "sensors/pb"
	"time"

	"github.com/labstack/echo/v4"
)

func (r *Handler) Update(c echo.Context) error {
	var (
		response model.Response
		payload  model.UpdateDataPayload
		in       sensorpb.PayloadUpdateSensors
	)

	err := c.Bind(&payload)
	if err != nil {
		log.Println(err.Error())
		response.Message = "Failed to bind payload"
		return c.JSON(http.StatusBadRequest, response)
	}

	in = sensorpb.PayloadUpdateSensors{
		Id_1:      payload.Id_1,
		Id_2:      payload.Id_2,
		Type:      payload.Type,
		Value:     payload.Value,
		Timestamp: payload.Timestamp,
	}

	if payload.From != nil && payload.To != nil {
		in.From = payload.From
		in.To = payload.To
	}

	if payload.FilterId1 != nil && payload.FilterId2 != nil {
		in.FilterId1 = payload.FilterId1
		in.FilterId2 = payload.FilterId2
	}

	datas, err := r.sensorGRPC.EditSensor(context.Background(), &in)
	if err != nil {
		log.Println(err.Error())
		response.Message = "Failed to edit sensor"
		return c.JSON(http.StatusBadRequest, response)
	}

	response = model.Response{
		StatusCode: http.StatusOK,
		Message:    "success",
		Status:     "OK",
		Timestamp:  time.Now().Format(time.RFC3339),
		Data:       datas,
	}

	return c.JSON(http.StatusOK, response)
}
