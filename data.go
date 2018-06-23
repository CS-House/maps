package main

import (
	"database/sql"
	"fmt"

	"github.com/Jeffail/gabs"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	input := `{"DeviceID":"867322035135813","TimeStamp":1527575284000,"Values":[{"Latitude":18.709738,"Longitude":80.068397,"Speed":0,"Box":false,"Battery":true,"Ignition":false}]}`

	jsonParsed, _ := gabs.ParseJSON([]byte(input))

	var DID string
	DID = jsonParsed.Path("DeviceID").Data().(string)
	// ts := jsonParsed.Path("TimeStamp")

	// fmt.Println(DID, ts)
	children := jsonParsed.Path("Values").Index(0)

	fmt.Println(children.Path("Latitude"))
	fmt.Println(children.Path("Longitude"))
	fmt.Println(children.Path("Speed"))
	fmt.Println(children.Path("Box"))
	fmt.Println(children.Path("Battery"))
	fmt.Println(children.Path("Ignition"))

	db, err := sql.Open("mysql", "root:clear@/latlong")

	stmtIns, err := db.Prepare("INSERT INTO GPSdata VALUES( ?, ?, ?, ?, ?, ?, ?, ?, ?)")

	//for i := 26; i < 50; i++ {
	_, err = stmtIns.Exec(DID, 1, 1, 1, 1, 1, 1, 1, 1)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	//}

	defer db.Close()

}

// {
// 	"DeviceID":"867322035135813",
// 	"TimeStamp":1527575284000,
// 	"values": [
// 		{
// 			"Latitude":18.709738,
// 			"Longitude":80.068397,
// 			"Speed":0,
// 			"Box":false,
// 			"Bat":true,
// 			"Ign":false
// 		}
// 	]
// }
