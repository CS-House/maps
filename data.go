package main

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

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
	fmt.Println(children.Path("Speed"), DID)

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
