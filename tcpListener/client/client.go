package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
)

const (
	server = "localhost:8000"
)

func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp", server)

	go handler(conn)

	fmt.Print("[CLIENT] Start talking with the server: ")

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
			fmt.Println("[CLIENT] ERROR: something wrong occured:", err.Error())
			return
		}
		n := bytes.Index(buf, []byte{0})
		fmt.Print("[CLIENT] echo: " + string(buf[:n]))
	}
}
