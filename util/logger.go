package util

import (
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

// Logger is the configured logging struct
var Logger *log.Logger

// LogFile is the file the log is written to
var LogFile os.File

func init() {

	// Log to stdout and a file
	LogFile, err := os.Create("./log/go-website-" + strconv.Itoa(int(time.Now().Unix())) + ".log")
	if err != nil {
		log.Fatal(err)
	}

	LogFile.Sync()
	Logger = log.New(io.MultiWriter(LogFile, os.Stdout), "", log.Ldate|log.Ltime)
}
