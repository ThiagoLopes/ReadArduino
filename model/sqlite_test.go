package model

import (
	"os"
	"testing"
	"time"
	"reflect"
)

const dbtest = "test.db"

func TestAll(t *testing.T) {
	db := InitDB(dbtest)
	defer db.Close()
	defer os.Remove(dbtest)

	CreateTable(db)

	nt := time.Now().UTC()

	serialdatas := []SerialData{
		SerialData{0, 100, 200, 300, 400, 500, nt},
		SerialData{0, 101, 202, 303, 404, 505, nt},
	}

	Insert(db, serialdatas)

	readItems := Read(db)
	if reflect.DeepEqual(readItems, serialdatas){
		t.Error("readItems is nos equal serialdatas")
	}
}
