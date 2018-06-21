package main

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

// type Export struct {
// 	DeviceID  string `json:"DeviceID"`
// 	TimeStamp string `json:"TimeStamp"`
// 	Values    []Values
// }

// type Values struct {
// 	Latitude  string `json:"Latitude"`
// 	Longitude string `json:"Longitude"`
// 	Speed     string `json:"Speed"`
// 	Box       string `json:"Box"`
// 	Battery   string `json:"Battery"`
// 	Ignition  string `json:"Ignition"`
// }

func main() {

	input := `{"DeviceID":"867322035135813","TimeStamp":1527575284000,"Values":[{"Latitude":18.709738,"Longitude":80.068397,"Speed":0,"Box":false,"Battery":true,"Ignition":false}]}`

	jsonParsed, _ := gabs.ParseJSON([]byte(input))

	children := jsonParsed.Path("Values").Index(0)

	fmt.Println(children)
	// rand.Seed(time.Now().UnixNano())

	// pubnub := messaging.NewPubnub(
	// 	pubkey,
	// 	subkey,
	// 	"",
	// 	"",
	// 	false,
	// 	"",
	// 	nil)

	// successChannel := make(chan []byte, 0)
	// errorChannel := make(chan []byte, 0)

	//message := "hello world from go publisher"

	// for {

	// 	// json, _ := json.Marshal()
	// 	// fmt.Println(string(json))

	// 	pubnub.Publish(
	// 		"exp-channel",
	// 		input,
	// 		successChannel,
	// 		errorChannel)

	// 	select {
	// 	case response := <-successChannel:
	// 		fmt.Println(string(response))
	// 	case err := <-errorChannel:
	// 		fmt.Println(string(err))
	// 	case <-messaging.Timeout():
	// 		fmt.Println("Publish() timeout")
	// 	}

	// 	time.Sleep(2 * time.Second)
	// }
}

// func floatToString(input float64) string {
// 	return strconv.FormatFloat(input, 'f', 2, 64)
// }

// const (
// 	pubkey = "pub-c-f3cae627-a107-45d2-a3cc-256467b09e6a"
// 	subkey = "sub-c-18580a92-f8cc-11e5-9086-02ee2ddab7fe"
// )

// type Packet struct {
// 	Latitude  string `json:"Latitude"`
// 	Longitude string `json:"Longitude"`
// 	DeviceID  string `json:"DeviceID"`
// }

// func (p *Packet) GetRandomLatAndLong(min, max float64) {
// 	p.Latitude = floatToString(min + rand.Float64()*(max-min))
// 	p.Longitude = floatToString(min + rand.Float64()*(max-min))
// }

// var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// func randSeq(n int) string {
// 	b := make([]rune, n)
// 	for i := range b {
// 		b[i] = letters[rand.Intn(len(letters))]
// 	}
// 	return string(b)
// }
