package main

import (
	"fmt"
	"github.com/ThiagoLopes/noir-client/config"
	"github.com/ThiagoLopes/noir-client/model"
	mserial "github.com/ThiagoLopes/noir-client/serial"
	"github.com/tarm/serial"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	TOKEN       = []byte("1")
	BAUD, _     = strconv.Atoi(config.GetEnvDefault("BAUD", "9600"))
	DATABASE    = config.GetEnvDefault("DATABASE", "serial.db")
	SERIAL_PATH = config.GetEnvDefault("ARDUINO", "/dev/ttyUSB0")
)

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
	go mserial.LoopReadAndPost(db, client)

	fmt.Scanln() // just dont close pls
	fmt.Println(model.Read(db, false))
}
