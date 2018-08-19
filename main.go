package main

import (
	"fmt"
	mserial "github.com/ThiagoLopes/ReadArduino/serial"
	"github.com/tarm/serial"
	"log"
	"os"
	"time"
)

var BUFFER_READ [][]byte
var TOKEN = []byte("1")
var SERIAL_PATH string

const BAUD = 9600

func init() {
	if env := os.Getenv("ARDUINO"); env == "" {
		SERIAL_PATH = "/dev/ttyUSB0"
	} else {
		SERIAL_PATH = env
	}
}

func main() {
	c := &serial.Config{
		Name:        SERIAL_PATH,
		Baud:        BAUD,
		ReadTimeout: time.Second / 2, // RUDE
	}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Panic(err)
	}
	s.Flush() // Clean data before start read
	go mserial.LoopWriteAndRead(s, &TOKEN, &BUFFER_READ)
	fmt.Scanln()
	fmt.Println(BUFFER_READ)
}
