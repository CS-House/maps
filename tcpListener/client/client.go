package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
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
		p := &Packet{}
		p.GetRandomLatAndLong(10, 100)

		//reader := bufio.NewReader(os.Stdin)
		//text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, "%s %s", p.Latitude, p.Longitude)

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
		// n := bytes.Index(buf, []byte{0})
		// fmt.Print("[CLIENT] echo: " + string(buf[:n]))
	}
}

type Packet struct {
	Latitude  string `json:"Latitude"`
	Longitude string `json:"Longitude"`
}

func (p *Packet) GetRandomLatAndLong(min, max float64) {
	p.Latitude = floatToString(min + rand.Float64()*(max-min))
	p.Longitude = floatToString(min + rand.Float64()*(max-min))
}

func floatToString(input float64) string {
	return strconv.FormatFloat(input, 'f', 2, 64)
}
