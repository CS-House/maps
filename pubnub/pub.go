package main

import (
	"fmt"

	"github.com/pubnub/go/messaging"
)

const (
	pubKey = "pub-c-f3cae627-a107-45d2-a3cc-256467b09e6a"
	//subKey    = "sub-c-18580a92-f8cc-11e5-9086-02ee2ddab7fe"
	//secretKey = "sec-c-MGMxMDBmYTMtOGZmYy00ZjMyLTkwMmUtMjE4YWE5MzJiYjg5"
)

func main() {
	fmt.Println("PubNub SDK for go - ", messaging.VersionInfo())

	pubnub := messaging.NewPubnub(pubKey, "", "", "", false, "", nil)

	successChannel := make(chan []byte)
	errorChannel := make(chan []byte)

	message := struct {
		First  string `json:"first"`
		Second string `json:"second"`
		Age    int    `json:"age"`
		Region string `json:"region"`
	}{
		"Robert",
		"Plant",
		59,
		"UK",
	}
	go pubnub.Publish(
		"hello",
		message,
		successChannel,
		errorChannel,
	)
	select {
	case response := <-successChannel:
		fmt.Println(string(response))
	case err := <-errorChannel:
		fmt.Println(string(err))
	case <-messaging.Timeout():
		fmt.Println("Publish() timeout")
	}
}
