package pubnub

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/pubnub/go/messaging"
)

func main() {

	// input := "GTPL $1,867322035135813,A,290518,062804,18.709738,N,80.068397,E,0,406,309,11,0,14,1,0,26.4470#"

	// input2 := "*ZJ,2030295125,V1,073614,A,3127.7080,N,7701.8360,E,0.00,0.00,040618,00000000#"

	// fmt.Println(parsepub.Parse(&input2))
	// fmt.Println(parsepub.Parse(&input))

	//o1 := parsepub.Parse(&input)

	rand.Seed(time.Now().UnixNano())

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

	//message := "hello world from go publisher"

	points := make([]string, 0)
	for index := 0; index < 5; index++ {
		p := randSeq(10)
		points = append(points, p)
	}

	for {

		p := &Packet{}
		p.GetRandomLatAndLong(10, 50)

		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(points))
		fmt.Println("index is ", index)
		p.DeviceID = points[index]
		json, _ := json.Marshal(p)
		fmt.Println(string(json))

		pubnub.Publish(
			"exp-channel",
			string(json),
			successChannel,
			errorChannel)

		time.Sleep(2 * time.Second)

		select {
		case response := <-successChannel:
			fmt.Println(string(response))
		case err := <-errorChannel:
			fmt.Println(string(err))
		case <-messaging.Timeout():
			fmt.Println("Publish() timeout")
		}
	}
}

func floatToString(input float64) string {
	return strconv.FormatFloat(input, 'f', 2, 64)
}

const (
	pubkey = "pub-c-f3cae627-a107-45d2-a3cc-256467b09e6a"
	subkey = "sub-c-18580a92-f8cc-11e5-9086-02ee2ddab7fe"
)

type Packet struct {
	Latitude  string `json:"Latitude"`
	Longitude string `json:"Longitude"`
	DeviceID  string `json:"DeviceID"`
}

func (p *Packet) GetRandomLatAndLong(min, max float64) {
	p.Latitude = floatToString(min + rand.Float64()*(max-min))
	p.Longitude = floatToString(min + rand.Float64()*(max-min))
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
