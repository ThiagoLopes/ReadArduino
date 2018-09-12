package model

import (
	"os"
	"reflect"
	"testing"
	"time"
	"net/http"
)

const dbtest = "test.db"

func TestInserAndRead(t *testing.T) {
	db := InitDB(dbtest)
	defer db.Close()
	defer os.Remove(dbtest)

	CreateTable(db)

	nt := time.Now().UTC()

	t.Run("TestInsertAndRead", func(t *testing.T) {
		serialdatas := []SerialData{
			SerialData{0, 100, 200, 300, 400, 500, nt},
			SerialData{0, 101, 202, 303, 404, 505, nt},
		}

		Insert(db, serialdatas)

		readItems := Read(db)
		if reflect.DeepEqual(readItems, serialdatas) {
			t.Error("readItems is nos equal serialdatas")
		}
	})

	t.Run("TestPostOrSaveDB", func(t *testing.T) {
		test_data := []byte("33.3,44.33,444.44,33.3,55.55")
		client := &http.Client{} // mock this
		PostOrSaveDB(test_data, db, client, "http://localhost:0000")
		t.Log(Read(db))
	})
}
