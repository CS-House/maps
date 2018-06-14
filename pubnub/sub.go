package main

import (
	"encoding/json"
	"fmt"

	"github.com/pubnub/go/messaging"
)

const (
	//pubKey    = "pub-c-f3cae627-a107-45d2-a3cc-256467b09e6a"
	subKey = "sub-c-18580a92-f8cc-11e5-9086-02ee2ddab7fe"
	//secretKey = "sec-c-MGMxMDBmYTMtOGZmYy00ZjMyLTkwMmUtMjE4YWE5MzJiYjg5"
)

func main() {
	fmt.Println("PubNub SDK for go - ", messaging.VersionInfo())

	pubnub := messaging.NewPubnub("", subKey, "", "", false, "", nil)

	successChannel := make(chan []byte)
	errorChannel := make(chan []byte)

	go pubnub.Subscribe("hello", "", successChannel, false, errorChannel)

	for {
		select {
		case response := <-successChannel:
			var msg []interface{}

			err := json.Unmarshal(response, &msg)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Received message '", response, "'")

			switch m := msg[0].(type) {
			case float64:
				fmt.Println(msg[1].(string))
			case []interface{}:
				fmt.Printf("Received message '%s' on channel '%s'\n", m[0], msg[2])
				return
			default:
				panic(fmt.Sprintf("Unknown type: %T", m))
			}

		case err := <-errorChannel:
			fmt.Println(string(err))
		case <-messaging.SubscribeTimeout():
			fmt.Println("Subscribe() timeout")
		}
	}
}
