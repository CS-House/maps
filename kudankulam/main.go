package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	db, err := sql.Open("mysql", "root:clear@/kudankulam")

	check(err)
	fmt.Println("db connected!")

	rows, err := db.Query("select distinct device_id from location_history_current")
	fmt.Println("rows taken!")
	var did string

	check(err)

	defer rows.Close()
	fmt.Println("deferred.")

	for rows.Next() {
		err := rows.Scan(&did)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(did)
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
