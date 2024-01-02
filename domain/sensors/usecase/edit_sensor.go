package usecase

import (
	"log"
	"sensors/domain/sensors/model"
	sensorpb "sensors/pb"
	"time"
)

func (s *service) EditSensors(in *sensorpb.PayloadUpdateSensors) (err error) {
	var (
		filter  model.FilterData
		payload model.SensorData
		layout  = "2006-01-02"
	)

	if in.From != nil && in.To != nil {
		// convert time into string
		from, err := time.Parse(layout, *in.From)
		if err != nil {
			log.Println(err.Error())
			return err
		}

		to, err := time.Parse(layout, *in.To)
		if err != nil {
			log.Println(err.Error())
			return err
		}

		filter.From = &from
		filter.To = &to
	}

	if in.FilterId1 != nil && in.FilterId2 != nil {
		filter.ID1 = in.FilterId1
		filter.ID2 = in.FilterId2
	}

	timestamp, err := time.Parse(layout, in.Timestamp)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	payload = model.SensorData{
		ID1:       in.Id_1,
		ID2:       int(in.Id_2),
		Value:     float64(in.Value),
		Type:      in.Type,
		Timestamp: timestamp,
	}

	err = s.sensorRepo.Update(payload, filter)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
