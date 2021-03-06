package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"

	"github.com/buger/jsonparser"
	_ "github.com/go-sql-driver/mysql"
	logger "github.com/gowtham-munukutla/maps/logger"
	"github.com/gowtham-munukutla/maps/parsepub"
	"github.com/pubnub/go/messaging"
)

const (
	port = "8000"
)

var db *sql.DB

var clients []net.Conn
var count = 0

func main() {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Ls.Fatal(err)
	}

	defer ln.Close()
	go signalHandler()

	log.Println("[SERVER] listening...")
	logger.Ls.Print("[SERVER] listening...")

	for {
		conn, err := ln.Accept()
		count++
		if err != nil {
			log.Println(err)
			logger.Ls.Print(err)
		}

		log.Printf("[SERVER] Client connected %s -> %s -- Number of clients connected (%d)\n", conn.RemoteAddr(), conn.LocalAddr(), count)
		logger.Ls.Printf("[SERVER] Client connected %s -> %s -- Number of clients connected (%d)\n", conn.RemoteAddr(), conn.LocalAddr(), count)
		// Add the client to the connection array
		clients = append(clients, conn)

		go handler(conn)
	}
}

func removeClient(conn net.Conn) {
	log.Printf("[SERVER] Client %s disconnected", conn.RemoteAddr())
	logger.Ls.Printf("[SERVER] Client %s disconnected", conn.RemoteAddr())
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

	db, err := sql.Open("mysql", "root:Digitalmysql@123@tcp(139.59.70.228:3306)/gowthammedigada")

	check(err)

	for {
		select {
		case data := <-dataChan:
			
			log.Printf("[SERVER} Client %s sent: %s", conn.RemoteAddr(), string(data))
			jsonObj := parsepub.Parse(string(data))

			log.Printf("[SERVER] Client %s sent converted json: %s", conn.RemoteAddr(), jsonObj)
			logger.Ls.Printf("[SERVER] Client %s sent: %s", conn.RemoteAddr(), jsonObj)

			var wg sync.WaitGroup
			wg.Add(2)

			go insertDB([]byte(jsonObj), db, &wg)

			go func(obj string, pnSuccChan chan []byte, pnErrorChan chan []byte, wg *sync.WaitGroup) {
				defer wg.Done()
				Pubnub.Publish(
					"exp-channel",
					obj,
					pnSuccChan,
					pnErrorChan)
			}(jsonObj, successChannel, errorChannel, &wg)

			wg.Wait()

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
			logger.Ls.Println("[SERVER] An error occured:", err.Error())
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
			logger.Ls.Printf("[SERVER] Closing due to Signal: %s", sig)
			log.Printf("[SERVER] Graceful shutdown")
			logger.Ls.Printf("[SERVER] Graceful shutdown")
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

func insertDB(object []byte, db *sql.DB, wg *sync.WaitGroup) {

	defer wg.Done()

	DeviceID, _, _, _ := jsonparser.Get(object, "DeviceID")
	TimeStamp, _, _, _ := jsonparser.Get(object, "TimeStamp")
	Latitude, _, _, _ := jsonparser.Get(object, "Latitude")
	Longitude, _, _, _ := jsonparser.Get(object, "Longitude")
	Battery, _, _, _ := jsonparser.Get(object, "Battery")
	Ignition, _, _, _ := jsonparser.Get(object, "Ignition")
	Speed, _, _, _ := jsonparser.Get(object, "Speed")
	Box, _, _, _ := jsonparser.Get(object, "Box")
	Alert, _, _, _ := jsonparser.Get(object, "Alert")

	_, err := db.Exec("Insert into GPSdata Values(?,?,?,?,?,?,?,?, ?)",
		string(DeviceID), string(TimeStamp), string(Latitude), string(Longitude), string(Speed), string(Box), string(Battery), string(Ignition), string(Alert))

	check(err)

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
