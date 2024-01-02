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

func (r *Handler) Get(c echo.Context) error {
	var (
		response model.Response
		payload  model.FilterData
		from, to string
		in       sensorpb.PayloadGetSensors
	)

	err := c.Bind(&payload)
	if err != nil {
		log.Println(err.Error())
		response.Message = "Failed to bind payload"
		return c.JSON(http.StatusBadRequest, response)
	}

	if payload.From != nil && payload.To != nil {
		from = payload.From.Format("2006-01-02")
		to = payload.To.Format("2006-01-02")

		in.From = &from
		in.To = &to
	}

	if payload.ID1 != nil && payload.ID2 != nil {
		in.Id_1 = payload.ID1
		in.Id_2 = payload.ID2
	}

	datas, err := r.sensorGRPC.GetSensors(context.Background(), &in)
	if err != nil {
		log.Println(err.Error())
		response.Message = "Failed to get sensor"
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
