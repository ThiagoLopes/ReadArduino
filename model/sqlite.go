package model

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const (
	DATE = "2006-01-02"
	TIME = "15:04:05 -0700 MST"
)

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Panic(err)
	}
	if db == nil {
		log.Panic("db nil")
	}
	return db
}

func CreateTable(db *sql.DB) {
	sql_table := `
	CREATE TABLE IF NOT EXISTS serialdata(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		humidity FLOAT,
		temperature FLOAT,
		co FLOAT,
		co2 FLOAT,
		mp25 FLOAT,
		created_date_at DATE,
		created_time_at TIME
	);
	`
	// add index

	_, err := db.Exec(sql_table)
	if err != nil {
		log.Panic(err)
	}
}

func Insert(db *sql.DB, sds []SerialData) {
	sql_addserial := `
	INSERT OR REPLACE INTO serialdata(
		humidity,
		temperature,
		co,
		co2,
		mp25,
		created_date_at,
		created_time_at,
	) values(?, ?, ?, ?, ?, ?, ?)
	`
	stmt, err := db.Prepare(sql_addserial)
	if err != nil {
		log.Panic(err)
	}
	defer stmt.Close()

	for _, sd := range sds {
		_, err := stmt.Exec(
			sd.Humidity,
			sd.Temperature,
			sd.CO,
			sd.CO2,
			sd.MP25,
			sd.CreatedDateAt.Format(TIME),
			sd.CreatedTimeAt.Format(DATE),
		)
		if err != nil {
			log.Panic(err)
		}
	}
}

func Read(db *sql.DB) []SerialData {
	sql_readall := `
	SELECT
		humidity,
		temperature,
		co,
		co2,
		mp25,
		created_date_at,
		created_time_at,
	FROM serialdata
	`

	rows, err := db.Query(sql_readall)
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()

	var result []SerialData
	for rows.Next() {
		sd := SerialData{}
		err = rows.Scan(
			&sd.Id,
			&sd.Humidity,
			&sd.Temperature,
			&sd.CO,
			&sd.CO2,
			&sd.MP25,
			&sd.CreatedDateAt,
			&sd.CreatedTimeAt,
		)
		if err != nil {
			log.Panic(err)
		}
		result = append(result, sd)
	}
	return result
}
