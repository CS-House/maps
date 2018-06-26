package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	dids := [10]string{"EUgnN1N1Eu", "MRl06xqlYR", "R6E2YeGHi5",
		"G9ceDX6fTT", "Zwkpbm93vj", "gFGKPNdTiO",
		"Py5qQAcf0i", "GlwuEiwrHO", "9WMPR3n0J9", "QkypngR4p4"}

	lats := [10]float64{85.23645, 23.45668, 36.12485, 89.45218, 62.31548, 49.56321, 78.13698, 49.65217, 36.91254, 69.13547}

	longs := [10]float64{62.15477, 76.31485, 53.12968, 79.13654, 49.36574, 43.85412, 19.21547, 63.28745, 37.26845, 95.63247}

	ln, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatalln("connection messed up", err.Error())
		panic(err)
	}
	defer ln.Close()

	for {
		var buffer bytes.Buffer

		str, integer := random(1, 9)
		buffer.WriteString("GTPL $" + str + "," + dids[integer] + ",A,")

		out, _ := exec.Command("/bin/bash", "-c", "date +%d%m%y,%I%M%S | tr -d '\n'").Output()

		lats[integer] = lats[integer] + 0.015
		latstr := FloatToString(lats[integer])

		buffer.WriteString(string(out) + "," + latstr + ",N,")

		longs[integer] = longs[integer] + 0.015
		longstr := FloatToString(longs[integer])
		buffer.WriteString(longstr + ",E,0,406,309,11,0,14,1,0,26.4470#")

		fmt.Println(buffer.String())

		m, err := ln.Write(buffer.Bytes())
		if err != nil {
			log.Fatalln("write messed up", err.Error())
			panic(err)
		}
		fmt.Println("Wrote ", strconv.Itoa(m)+", bytes")

		time.Sleep(2 * time.Second)
	}
}

func random(min, max int) (string, int) {
	rand.Seed(time.Now().Unix())
	integer := rand.Intn(max-min) + min

	return strconv.Itoa(integer), integer
}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 5, 32)
}

// GTPL $1,867322035135813,A,290518,062804,18.709738,N,80.068397,E,0,406,309,11,0,14,1,0,26.4470#
// $GPRMC,002832.000,A,1843.0858,N,08004.7547,E,0000,314.63,260618,,,A*
