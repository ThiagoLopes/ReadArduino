package model

import (
	"os"
	"testing"
	"time"
)

const dbtest = "test.db"

func TestAll(t *testing.T) {
	db := InitDB(dbtest)
	defer db.Close()
	defer os.Remove(dbtest)

	CreateTable(db)

	nt := time.Now().UTC()

	serialdatas := []SerialData{
		SerialData{0, 100, 200, 300, 400, 500, nt, nt},
		SerialData{0, 101, 202, 303, 404, 505, nt, nt},
	}

	Insert(db, serialdatas)

	readItems := Read(db)
	t.Log(readItems)
}
