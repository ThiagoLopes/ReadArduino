package main

import (
	"fmt"
	"github.com/ThiagoLopes/noir-client/model"
	mserial "github.com/ThiagoLopes/noir-client/serial"
	"github.com/tarm/serial"
	"log"
	"net/http"
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
	client := &http.Client{}

	s.Flush() // Clean data before start read

	go mserial.LoopWriteReadAndSave(s, &TOKEN, db, client) //start write, read and save

	fmt.Scanln() // just dont close pls
	fmt.Println(model.Read(db))
}
