package model

import (
	sensorpb "sensors/pb"
	"time"
)

type SensorData struct {
	Value     float64   `json:"value"`
	Type      string    `json:"type"`
	ID1       string    `json:"id_1"` // alphabets in capital letter
	ID2       int       `json:"id_2"`
	Timestamp time.Time `json:"timestamp"`
}

type FilterData struct {
	ID1  *string    `json:"id_1"` // alphabets in capital letter
	ID2  *int32     `json:"id_2"`
	From *time.Time `json:"from"`
	To   *time.Time `json:"to"`
}

func (s SensorData) ToPB() *sensorpb.Sensor {
	return &sensorpb.Sensor{
		Id_1:      s.ID1,
		Id_2:      int32(s.ID2),
		Type:      s.Type,
		Value:     float32(s.Value),
		Timestamp: s.Timestamp.Format("2006-01-02"),
	}
}
