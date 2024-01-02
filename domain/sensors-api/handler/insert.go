package handler

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"sensors/domain/sensors-api/model"
	sensorpb "sensors/pb"
	"time"

	"github.com/labstack/echo/v4"
)

func (r *Handler) Insert(c echo.Context) error {
	var response model.Response

	in := sensorpb.PayloadRequest{
		Id_1:      toCharStr(),
		Id_2:      int32(rand.Intn(100-1) + 1),
		Type:      "sensor",
		Value:     float32(1.00 + rand.Float64()*(100.0-1.0)),
		Timestamp: time.Now().Format("2006-01-02"),
	}

	_, err := r.sensorGRPC.InsertSensor(context.Background(), &in)
	if err != nil {
		log.Println(err.Error())
		response.Message = "Failed to generate sensor"
		return c.JSON(http.StatusBadRequest, response)
	}

	response = model.Response{
		StatusCode: http.StatusOK,
		Message:    "success",
		Status:     "OK",
		Timestamp:  time.Now().Format(time.RFC3339),
	}

	return c.JSON(http.StatusOK, response)
}

func toCharStr() string {
	i := rand.Intn(26-1) + 1

	return string('A' - 1 + i)
}
