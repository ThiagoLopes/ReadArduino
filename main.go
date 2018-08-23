package main

import (
	"fmt"
	mserial "github.com/ThiagoLopes/ReadArduino/serial"
	"github.com/ThiagoLopes/ReadArduino/model"
	"github.com/tarm/serial"
	"log"
	"os"
	"time"
)

var TOKEN = []byte("1")
var SERIAL_PATH string

const BAUD = 9600
const DATABASE = "serial.db"

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

	db := model.InitDB(DATABASE) // init db
	model.CreateTable(db)
	defer db.Close()

	s, err := serial.OpenPort(c)
	if err != nil {
		log.Panic(err)
	}

	s.Flush() // Clean data before start read

	go mserial.LoopWriteAndReadAndSave(s, &TOKEN, db) //start write, read and save

	fmt.Scanln() // just dont close pls
	fmt.Println(model.Read(db))
}
