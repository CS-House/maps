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
	lcpath = lcpath + "/src/github.com/gowtham-munukutla/maps/logs/client.log"
	flag.Parse()

	file, err := os.OpenFile(lcpath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	Lc = log.New(file, "", log.LstdFlags|log.Lshortfile)
}
