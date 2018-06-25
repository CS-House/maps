package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/gowtham-munukutla/maps/logger"
)

const (
	server = "localhost:5000"
)

func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp", server)

	go handler(conn)

	logger.Lc.Print("[CLIENT] Start talking with the server: ")
	log.Println("[CLIENT] Start talking with the server: ")

	go signalHandler()

	// go readLine(conn)
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		// messagechannel := make(chan string, 500)
		// readLine("../../data.txt", messagechannel)

		// msg := receive(messagechannel)
		// fmt.Println(msg)

		// send to socket
		fmt.Fprintf(conn, text)
	}
}

func handler(conn net.Conn) {
	for {
		// listen for reply
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			log.Println("[CLIENT] ERROR: Maybe the server went offline: Please check : ")
			logger.Lc.Println("[CLIENT] ERROR: Maybe the server went offline: Please check : ", err.Error())
			return
		}
		// n := bytes.Index(buf, []byte{0})
		// fmt.Print("[CLIENT] echo: " + string(buf[:n]))
	}
}

func signalHandler() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	go func() {
		for sig := range sigchan {
			log.Printf("[CLIENT] Disconnecting: %s", sig)
			logger.Lc.Printf("[CLIENT] Disconnecting: %s", sig)
			log.Println("Done.")
			logger.Lc.Println("Done.")
			// Exit cleanly
			os.Exit(0)
		}
	}()
}
