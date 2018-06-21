package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
)

const (
	server = "localhost:8000"
)

func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp", server)

	go handler(conn)

	fmt.Print("[CLIENT] Start talking with the server: ")

	go signalHandler()

	for {
		// read in input from stdin

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
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
			fmt.Println("[CLIENT] ERROR: Maybe the server went offline: Please check : ", err.Error())
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
			fmt.Println("Done.")
			// Exit cleanly
			os.Exit(0)
		}
	}()
}
