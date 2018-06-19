package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
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
	defer removeClient(conn)
	errorChan := make(chan error)
	dataChan := make(chan []byte)

	go readWrapper(conn, dataChan, errorChan)

	for {
		select {
		case data := <-dataChan:
			log.Printf("[SERVER] Client %s sent: %s", conn.RemoteAddr(), string(data))
			for i := range clients {
				clients[i].Write(data)
			}
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
