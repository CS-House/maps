package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/pubnub/go/messaging"
)

func main() {
<<<<<<< HEAD

	rand.Seed(time.Now().UnixNano())

	p := &Packet{}

	p.GetRandomLatAndLong(10, 30)
=======
>>>>>>> e2b51198fe60f86eb9109dcfeb11cdbdc90a05d9

	json, _ := json.Marshal(p)

<<<<<<< HEAD
=======
	p := &Packet{}
	p.GetRandomLatAndLong(10, 50)

	json, _ := json.Marshal(p)

>>>>>>> e2b51198fe60f86eb9109dcfeb11cdbdc90a05d9
	fmt.Println(string(json))

	pubnub := messaging.NewPubnub(
		pubkey,
		subkey,
		"",
		"",
		false,
		"",
		nil)

	successChannel := make(chan []byte)
	errorChannel := make(chan []byte)

	//message := "hello world from go publisher"

	pubnub.Publish(
		"exp-channel",
		string(json),
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
}

<<<<<<< HEAD
func floatToString(input float64) string {
	return strconv.FormatFloat(input, 'f', 2, 64)
}

=======
>>>>>>> e2b51198fe60f86eb9109dcfeb11cdbdc90a05d9
const (
	pubkey = "pub-c-f3cae627-a107-45d2-a3cc-256467b09e6a"
	subkey = "sub-c-18580a92-f8cc-11e5-9086-02ee2ddab7fe"
)

type Packet struct {
	Latitude  string `json:"Latitude"`
	Longitude string `json:"Longitude"`
}

func (p *Packet) GetRandomLatAndLong(min, max float64) {
<<<<<<< HEAD
	p.Latitude = floatToString(min + rand.Float64()*(max-min))
	p.Longitude = floatToString(min + rand.Float64()*(max-min))
}

=======
	p.Latitude = strconv.FormatFloat(min+rand.Float64()*(max-min), 'f', 3, 64)
	p.Longitude = strconv.FormatFloat(min+rand.Float64()*(max-min), 'f', 3, 64)
}
>>>>>>> e2b51198fe60f86eb9109dcfeb11cdbdc90a05d9
