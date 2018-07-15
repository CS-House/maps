package main

import (
	"database/sql"
)

var db *sql.DB

func main() {
	db, err := sql.Open("mysql", "")

	check(err)

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type JsonObj struct {
	deviceID string
	lat      string
	long     string
}
