package util

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

var Logger logger
var LogFile os.File
var logFileBuffer *bufio.Writer
var stdErrBuffer *bufio.Writer

type logger struct {
	wrappedLogger *log.Logger
}

func init() {

	// Log to terminal and a file
	LogFile, err := os.Create("./log/go-website-" + strconv.Itoa(int(time.Now().Unix())) + ".log")
	if err != nil {
		log.Fatal(err)
	}

	LogFile.Sync()

	logFileBuffer = bufio.NewWriter(LogFile)
	stdErrBuffer = bufio.NewWriter(os.Stderr)
	Logger = logger{
		wrappedLogger: log.New(io.MultiWriter(logFileBuffer, stdErrBuffer), "", log.Ldate|log.Ltime),
	}
}

func (l *logger) Println(text interface{}) {
	l.wrappedLogger.Println(text)
	logFileBuffer.Flush()
	stdErrBuffer.Flush()
}

func (l *logger) Fatal(text interface{}) {
	l.wrappedLogger.Println(text)
	logFileBuffer.Flush()
	stdErrBuffer.Flush()
	os.Exit(1)
}
