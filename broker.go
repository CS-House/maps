package main

import (
	"database/sql"
)

var db *sql.DB

func main() {
	db, err := sql.Open("mysql", "rootone:Test!@#$@tcp(139.59.90.102:3306)/kudankalam")

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
