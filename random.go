// // script to generate random data for testing
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	"time"
// )

// func main() {
// 	//input := `{"DeviceID":"867322035135813","TimeStamp":1527575284000,"Values":[{"Latitude":18.709738,"Longitude":80.068397,"Speed":0,"Box":false,"Battery":true,"Ignition":false}]}`

// 	points := make([]string, 0)
// 	for index := 0; index < 10; index++ {
// 		p := randSeq(10)
// 		points = append(points, p)
// 	}

// 	rand.Seed(time.Now().UnixNano())

// 	ar := make([]*ExportPacket, 0)

// 	for i := 0; i < 1000; i++ {
// 		e1 := &ExportPacket{}
// 		index := rand.Intn(len(points))
// 		e1.DeviceID = points[index]
// 		e1.LatAndLong(10, 100)
// 		ar = append(ar, e1)
// 	}

// 	f, _ := os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE, 0755)

// 	for i := 0; i < 500; i++ {
// 		b, _ := json.Marshal(ar[i])
// 		fmt.Fprintf(f, "`"+string(b)+"`\n")
// 	}
// }

// type ExportPacket struct {
// 	DeviceID  string `json:"DeviceID"`
// 	TimeStamp string `json:"TimeStamp`
// 	Latitude  string `json:"Latitude"`
// 	Longitude string `json:"Longitude"`
// 	Speed     string `json:"Speed"`
// 	Box       string `json:"Box"`
// 	Battery   string `json:"Battery"`
// 	Ignition  string `json:"Ignition"`
// }

// func (e *ExportPacket) LatAndLong(min, max float64) {
// 	e.TimeStamp = "1527575284000"
// 	e.Latitude = floatToString(min + rand.Float64()*(max-min))
// 	e.Longitude = floatToString(min + rand.Float64()*(max-min))
// 	e.Ignition = "dfdf"
// 	e.Battery = "sdve"
// 	e.Box = "vdfv"
// 	e.Speed = "srgver"
// }

// func floatToString(input float64) string {
// 	return strconv.FormatFloat(input, 'f', 2, 64)
// }

// var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// func randSeq(n int) string {
// 	b := make([]rune, n)
// 	for i := range b {
// 		b[i] = letters[rand.Intn(len(letters))]
// 	}
// 	return string(b)
// }

// // {
// // 	"DeviceID":"867322035135813",
// // 	"TimeStamp":1527575284000,
// // 	"values": [
// // 		{
// // 			"Latitude":18.709738,
// // 			"Longitude":80.068397,
// // 			"Speed":0,
// // 			"Box":false,
// // 			"Bat":true,
// // 			"Ign":false
// // 		}
// // 	]
// // }

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	messagechannel := make(chan string, 10)
	readLine("data.txt", messagechannel)

	msg := receive(messagechannel)
	fmt.Println(msg)
}

func readLine(path string, message chan<- string) {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		message <- scanner.Text()
		time.Sleep(1 * time.Second)
	}
}

func receive(ch <-chan string) string {
	return <-ch
}
