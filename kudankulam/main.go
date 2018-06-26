package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	db, err := sql.Open("mysql", "rootone:Test!@#$@tcp(139.59.90.102:3306)/kudankulam")

	check(err)
	fmt.Println("db connected!")

	rows, err := db.Query("select lat_message,lon_message from location_history_current")
	fmt.Println("rows taken!")
	var lat, long string

	check(err)

	defer rows.Close()
	fmt.Println("deferred.")

	for rows.Next() {
		err := rows.Scan(&lat, &long)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(lat, long)
	}

	err = rows.Err()

	if err != nil {
		log.Fatal(err)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
