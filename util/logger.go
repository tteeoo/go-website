package util

import (
	"bufio"
	"io"
	"log"
	"os"
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
	LogFile, err := os.Create("./go-website.log")
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
	l.wrappedLogger.Fatal(text)
	logFileBuffer.Flush()
	stdErrBuffer.Flush()
}
