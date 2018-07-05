package main

import (
	"database/sql"
	"log"

	"encoding/json"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pubnub/go/messaging"
)

func main() {
	db, err := sql.Open("mysql", "rootone:Test!@#$@tcp(139.59.90.102:3306)/kudankulam")
	defer db.Close()

	check(err)

	pubnub := messaging.NewPubnub(
		pubkey,
		subkey,
		"",
		"",
		false,
		"",
		nil)

	successChannel := make(chan []byte, 0)
	errorChannel := make(chan []byte, 0)

	for {
		fmt.Println("\n")
		log.Println("---DEVICE_ID--- -LATITUDE-  LONGITUDE   ---TIME_CREATED---  SPEED")

		jsonArray := fetchLatest(db)

		jsonObj, _ := json.Marshal(jsonArray)

		// Publish to pubnub
		pubnub.Publish(
			"exp-channel",
			string(jsonObj),
			successChannel,
			errorChannel)

		select {
		case response := <-successChannel:
			fmt.Println(string(response))
		case err := <-errorChannel:
			fmt.Println(string(err))
		case <-messaging.Timeout():
			fmt.Println("Publish() timeout")
		}

		time.Sleep(4 * time.Second)
	}

}

func fetchLatest(db *sql.DB) []*SingleJson {
	rows, err := db.Query("select device_id, lat_message, lon_message, created_date,speed from location_history_current where device_id<>'' and (device_id, created_date) IN (select device_id, max(created_date) from location_history_current group by device_id);")
	var did,lat,lon,time,speed string

	check(err)

	defer rows.Close()

	var jsonArray = make([]*SingleJson, 0)

	for rows.Next() {
		err := rows.Scan(&did, &lat,&lon,&time,&speed)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(did,lat,lon,time,speed)
		jsonElement := &SingleJson{
			DeviceID: did,
			Lat:      lat,
			Long:     lon,
		}
		jsonArray = append(jsonArray, jsonElement)
	}
	return jsonArray
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type SingleJson struct {
	DeviceID string
	Lat      string
	Long     string
}

const (
	pubkey = "pub-c-f3cae627-a107-45d2-a3cc-256467b09e6a"
	subkey = "sub-c-18580a92-f8cc-11e5-9086-02ee2ddab7fe"
)
