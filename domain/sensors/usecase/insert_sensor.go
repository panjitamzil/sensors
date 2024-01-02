package usecase

import (
	"fmt"
	"log"
	"sensors/domain/sensors/model"
	sensorpb "sensors/pb"
	"time"
)

func (s *service) InsertSensor(in *sensorpb.PayloadRequest) error {
	// convert time into string
	layout := "2006-01-02"
	timestamp, err := time.Parse(layout, in.Timestamp)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	// mapping value
	payload := model.SensorData{
		ID1:       in.Id_1,
		ID2:       int(in.Id_2),
		Type:      in.Type,
		Value:     float64(in.Value),
		Timestamp: timestamp,
	}

	fmt.Println(payload)

	err = s.sensorRepo.Insert(payload)
	if err != nil {
		log.Println(234, err.Error())
		return err
	}

	return nil
}
