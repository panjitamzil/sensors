package usecase

import (
	"log"
	"sensors/domain/sensors/model"
	sensorpb "sensors/pb"
	"time"
)

func (s *service) DeleteSensors(in *sensorpb.PayloadGetSensors) (err error) {
	var (
		filter model.FilterData
		layout = "2006-01-02"
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

	if in.Id_1 != nil && in.Id_2 != nil {
		filter.ID1 = in.Id_1
		filter.ID2 = in.Id_2
	}

	err = s.sensorRepo.Delete(filter)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
