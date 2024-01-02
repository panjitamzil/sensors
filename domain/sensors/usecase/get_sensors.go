package usecase

import (
	"log"
	"sensors/domain/sensors/model"
	sensorpb "sensors/pb"
	"time"
)

func (s *service) GetSensors(in *sensorpb.PayloadGetSensors) (result []model.SensorData, err error) {
	var (
		payload model.FilterData
	)

	if in.From != nil && in.To != nil {
		// convert time into string
		layout := "2006-01-02"
		from, err := time.Parse(layout, *in.From)
		if err != nil {
			log.Println(err.Error())
			return result, err
		}

		to, err := time.Parse(layout, *in.To)
		if err != nil {
			log.Println(err.Error())
			return result, err
		}

		payload.From = &from
		payload.To = &to
	}

	if in.Id_1 != nil && in.Id_2 != nil {
		payload.ID1 = in.Id_1
		payload.ID2 = in.Id_2
	}

	result, err = s.sensorRepo.GetAll(&payload)
	if err != nil {
		log.Println(err.Error())
		return result, err
	}

	return result, nil
}
