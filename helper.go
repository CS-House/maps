package main

import (
	"fmt"

	"github.com/buger/jsonparser"
)

func main() {
	input := []byte(`{"DeviceID":"867322035135813","TimeStamp":"1527575284000","Latitude":18.709738,"Longitude":80.068397,"Speed":0,"Box":false,"Battery":true,"Ignition":false}`)

	insert(input)
}

func insert(object []byte) {
	DeviceID, _, _, _ := jsonparser.Get(object, "DeviceID")
	TimeStamp, _, _, _ := jsonparser.Get(object, "TimeStamp")
	Latitude, _, _, _ := jsonparser.Get(object, "Latitude")
	Longitude, _, _, _ := jsonparser.Get(object, "Longitude")
	Battery, _, _, _ := jsonparser.Get(object, "Battery")
	Ignition, _, _, _ := jsonparser.Get(object, "Ignition")

	fmt.Println(string(DeviceID), string(TimeStamp), string(Latitude), string(Longitude), string(Battery), string(Ignition))
}

//GTPL $2,867322035135825,A,290517,062804,18.709738,S,80.068397,W,0#
