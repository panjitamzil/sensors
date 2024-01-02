package main

import (
	"log"
	"net"
	"sensors/config"
	"sensors/domain/sensors/handler"
	"sensors/domain/sensors/repository"
	"sensors/domain/sensors/usecase"
	sensorpb "sensors/pb"

	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	db, err := config.InitDatabase(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sensorRepository := repository.NewRepo(db)
	sensorUsecase := usecase.NewService(usecase.Dependencies{
		SensorRepo: sensorRepository,
		Config:     cfg,
	})
	sensorHandler := handler.NewHandler(sensorUsecase)

	// Start gRPC server
	listen, err := net.Listen("tcp", ":"+cfg.GRPC.Port)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	sensorpb.RegisterSensorServicesServer(grpcServer, sensorHandler)

	log.Println("Server listening on :50051")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal(err)
	}

}
