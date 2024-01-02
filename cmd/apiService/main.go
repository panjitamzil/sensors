package main

import (
	"log"
	"sensors/config"
	"sensors/domain/sensors-api/handler"
	sensorpb "sensors/pb"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	grpcConn, err := grpc.Dial(":"+cfg.GRPC.Port, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	grpcService := sensorpb.NewSensorServicesClient(grpcConn)

	e := echo.New()

	h := handler.NewHandler(grpcService)
	e.POST("/generate", h.Insert)
	e.GET("/get", h.Get)
	e.PUT("/update", h.Update)
	e.DELETE("/delete", h.Delete)

	if err = e.Start(":" + cfg.Application.Port); err != nil {
		log.Fatal("Failed to start the server: ", err.Error())
	}

}
