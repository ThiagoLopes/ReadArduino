package serial

import (
	"database/sql"
	"github.com/ThiagoLopes/noir-client/model"
	"github.com/tarm/serial"
	"log"
	"time"
	"net/http"
)

const (
	TIME_WHEN_ERROR = 5 * time.Second
	MAX_LEN_MESSAGE = 60
	MSG_PER_TIME    = 1000 * time.Millisecond
	HOST = "http://localhost:8000"
)

func readSerialWithBuffer(s *serial.Port) []byte {
	buf_message := make([]byte, MAX_LEN_MESSAGE)
	n, err := s.Read(buf_message)
	if err != nil {
		log.Println(err)
		time.Sleep(TIME_WHEN_ERROR)
	}
	log.Println(string(buf_message[:n]))
	return []byte(buf_message[:n])
}

func writeSerialToken(s *serial.Port, token *[]byte) {
	_, err := s.Write(*token)
	if err != nil {
		log.Println(err)
	}
}

func LoopWriteReadAndSave(s *serial.Port, t *[]byte, db *sql.DB, c *http.Client) {
	for {
		writeSerialToken(s, t)
		response_bytes := readSerialWithBuffer(s) // implement a err here
		go model.PostOrSaveDB(response_bytes, db, c, HOST)
		time.Sleep(MSG_PER_TIME)
	}
}
