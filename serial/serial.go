package serial

import (
	"bytes"
	"errors"
	"github.com/tarm/serial"
	"log"
	"time"
)

const (
	TIME_WHEN_ERROR = 5 * time.Second
	MAX_LEN_MESSAGE = 60
	MSG_PER_TIME    = time.Second
)

func ReadSerialWithBuffer(s *serial.Port, b *[][]byte) {
	buf_message := make([]byte, MAX_LEN_MESSAGE)
	n, err := s.Read(buf_message)
	if err != nil {
		log.Fatal(err)
		time.Sleep(TIME_WHEN_ERROR)
	}
	*b = append(*b, buf_message[:n])
	log.Println(string(buf_message[:n]))

}

func LoopWriteAndRead(s *serial.Port, t *[]byte, b *[][]byte) {
	for {
		WriteSerialToken(s, t)
		ReadSerialWithBuffer(s, b)
		time.Sleep(MSG_PER_TIME)
	}
}

func WriteSerialToken(s *serial.Port, token *[]byte) {
	_, err := s.Write(*token)
	if err != nil {
		log.Fatal(err)
	}
}

func NormalizeMessage(b *[]byte, expected_size int) ([][]byte, error) {
	splited_message := bytes.Split(*b, []byte(","))
	if len(splited_message) == expected_size {
		return splited_message, nil
	}
	return splited_message, errors.New("Message received by the serial is not valid")
}
