package model

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"
)

type SerialData struct {
	Id          int       `json:"id"`
	Humidity    float64   `json:"humidity"`
	Temperature float64   `json:"temperature"`
	CO          float64   `json:"co"`
	CO2         float64   `json:"co2"`
	MP25        float64   `json:"mp25"`
	CreatedAt   time.Time `json:"created_at"`
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
			return nil, errors.New("Fail parse to Float64")
		}
	}
	return message, nil
}

func NewSerialData(m []byte) (*SerialData, error) {
	normalized, err := normalizeMessage(&m, 5)
	if err != nil {
		return nil, err
	}
	return &SerialData{
		0,
		float64(normalized[0]),
		float64(normalized[1]),
		float64(normalized[2]),
		float64(normalized[3]),
		float64(normalized[4]),
		time.Now().UTC(),
	}, nil
}

func (ss *SerialData) DecodeAndPost(c *http.Client, url string) (*http.Response, error) {
	data, err := json.Marshal(ss)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// create a func
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer request.Body.Close()
	request.Header.Set("Content-Type", "application/json")

	response, err := c.Do(request)
	if err != nil {
		log.Println(err)
		return response, err
	}

	log.Println("status: ", response.Status)
	return response, nil
}

func PostOrSaveDB(bytes_recive []byte, db *sql.DB, c *http.Client, url string) {
	serial_data, err := NewSerialData(bytes_recive)
	if err != nil {
		log.Println(err)
		return
	}

	response, err := serial_data.DecodeAndPost(c, url)
	if err != nil || response.StatusCode != 202 {
		Insert(db, []SerialData{*serial_data})
		return
	}
	log.Println("Successfully POST")
}
