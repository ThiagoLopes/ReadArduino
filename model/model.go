package model

import "time"

type SerialData struct {
	Id            int
	Humidity      float64
	Temperature   float64
	CO            float64
	CO2           float64
	MP25          float64
	CreatedDateAt time.Time
	CreatedTimeAt time.Time
}
