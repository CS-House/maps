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
	lspath = lspath + "/src/gowtham-munukutla/logs/server.log"
	flag.Parse()

	file, err := os.Create(lspath)
	if err != nil {
		panic(err)
	}

	Ls = log.New(file, "", log.LstdFlags|log.Lshortfile)
}
