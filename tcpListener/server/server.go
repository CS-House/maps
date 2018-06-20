package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/gowtham-munukutla/maps/parsepub"
	"github.com/pubnub/go/messaging"
)

const (
	port = "8000"
)

var clients []net.Conn
var count = 0

func main() {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()
	go signalHandler()

	log.Print("[SERVER] listening...")

	for {
		conn, err := ln.Accept()
		count++
		if err != nil {
			log.Print(err)
		}

		log.Printf("[SERVER] Client connected %s -> %s -- Number of clients connected (%d)\n", conn.RemoteAddr(), conn.LocalAddr(), count)
		// Add the client to the connection array
		clients = append(clients, conn)

		go handler(conn)
	}
}

func removeClient(conn net.Conn) {
	log.Printf("[SERVER] Client %s disconnected", conn.RemoteAddr())
	count--
	conn.Close()
	//remove client from clients here
}

func handler(conn net.Conn) {

	successChannel := make(chan []byte, 0)
	errorChannel := make(chan []byte, 0)

	defer removeClient(conn)
	errorChan := make(chan error)
	dataChan := make(chan []byte)

	go readWrapper(conn, dataChan, errorChan)

	Pubnub := messaging.NewPubnub(
		pubkey,
		subkey,
		"",
		"",
		false,
		"",
		nil)

	for {
		select {
		case data := <-dataChan:
			log.Printf("[SERVER] Client %s sent: %s", conn.RemoteAddr(), string(data))

			jsonObj := parsepub.Parse(string(data))

			go Pubnub.Publish(
				"exp-channel",
				jsonObj,
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

			// for i := range clients {
			// 	clients[i].Write(data)
			// }
		case err := <-errorChan:
			log.Println("[SERVER] An error occured:", err.Error())
			return
		}
	}
}

func readWrapper(conn net.Conn, dataChan chan []byte, errorChan chan error) {
	for {
		buf := make([]byte, 1024)
		reqLen, err := conn.Read(buf)
		if err != nil {
			errorChan <- err
			return
		}
		dataChan <- buf[:reqLen]
	}
}

func signalHandler() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	go func() {
		for sig := range sigchan {
			log.Printf("[SERVER] Closing due to Signal: %s", sig)
			log.Printf("[SERVER] Graceful shutdown")
			fmt.Println("Done.")
			// Exit cleanly
			os.Exit(0)
		}
	}()
}

const (
	pubkey = "pub-c-f3cae627-a107-45d2-a3cc-256467b09e6a"
	subkey = "sub-c-18580a92-f8cc-11e5-9086-02ee2ddab7fe"
)
