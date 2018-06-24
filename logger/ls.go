package logger

import (
	"flag"
	"log"
	"os"
)

//create new custom logger for server.
var (
	Ls *log.Logger
)

func init() {
	lspath := os.Getenv("GOPATH")
	lspath = lspath + "/src/github.com/gowtham-munukutla/maps/logs/server.log"
	flag.Parse()

	file, err := os.OpenFile(lspath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	Ls = log.New(file, "", log.LstdFlags|log.Lshortfile)
}
