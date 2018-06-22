package logger

import (
	"flag"
	"log"
	"os"
)

//create custom new logger for client.
var (
	Lc *log.Logger
)

func init() {
	lcpath := os.Getenv("GOPATH")
	lcpath = lcpath + "/src/gowtham-munukutla/logs/client.log"
	flag.Parse()

	file, err := os.Create(lcpath)
	if err != nil {
		panic(err)
	}

	Lc = log.New(file, "", log.LstdFlags|log.Lshortfile)
}
