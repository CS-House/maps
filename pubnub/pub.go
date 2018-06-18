package main

import (
	"fmt"

	"github.com/pubnub/go/messaging"
)

const (
	pubkey = "pub-c-f3cae627-a107-45d2-a3cc-256467b09e6a"
	subkey = "sub-c-18580a92-f8cc-11e5-9086-02ee2ddab7fe"
)

func main() {

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

	message := "hello world from go publisher"

	pubnub.Publish(
		"exp-channel",
		message,
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
