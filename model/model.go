package model

import (
	"bytes"
	"database/sql"
	"errors"
	"log"
	"strconv"
	"time"
)

type SerialData struct {
	Id          int
	Humidity    float64
	Temperature float64
	CO          float64
	CO2         float64
	MP25        float64
	CreatedAt   time.Time
}

func normalizeMessage(b *[]byte, expected_size int) ([]float64, error) {
	splited_bytes := bytes.Split(*b, []byte(","))
	if len(splited_bytes) != expected_size {
		return nil, errors.New("Message received by the serial is not valid")

	}
	message := make([]float64, expected_size)
	for i, value := range splited_bytes {
		var err error
		message[i], err = strconv.ParseFloat(string(value), 64)
		if err != nil {
			errors.New("Fail parse to Float64")
		}
	}
	return message, nil
}

func NewSerialData(m []byte) *SerialData {
	normalized, err := normalizeMessage(&m, 5)
	if err != nil {
		log.Println(err)
	}
	return &SerialData{
		0,
		float64(normalized[0]),
		float64(normalized[1]),
		float64(normalized[2]),
		float64(normalized[3]),
		float64(normalized[4]),
		time.Now().UTC(),
	}
}

func PostOrSaveDB(bytes_recive []byte, db *sql.DB) {
	// write post method, currently save only
	serial_data := NewSerialData(bytes_recive)
	Insert(db, []SerialData{*serial_data})
}
