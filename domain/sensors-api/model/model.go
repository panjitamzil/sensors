package model

import (
	sensorpb "sensors/pb"
	"time"
)

type SensorDataPayload struct {
	Value     float64   `json:"value"`
	Type      string    `json:"type"`
	ID1       string    `json:"id_1"` // alphabets in capital letter
	ID2       int       `json:"id_2"`
	Timestamp time.Time `json:"timestamp"`
}

type UpdateDataPayload struct {
	Id_1      string  `json:"id_1"`
	Id_2      int32   `json:"id_2"`
	Type      string  `json:"type"`
	Value     float32 `json:"value"`
	Timestamp string  `json:"timestamp"`
	FilterId1 *string `json:"filter_id1"`
	FilterId2 *int32  `json:"filter_id2"`
	From      *string `json:"from"`
	To        *string `json:"to"`
}

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

type Response struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Status     string      `json:"status"`
	Timestamp  string      `json:"timestamp"`
	Data       interface{} `json:"data"`
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
