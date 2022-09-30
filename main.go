package main

import (
	"io"
	"log"
	"os"
	"time"

	r "go-company-api/routes"
)

const (
	Ldate     = 1 << iota // the date in the local time zone: 2009/01/23
	YYYYMMDD  = "2006-01-02"
	HHMMSS12h = "3:04:05 PM"
)

func main() {
	setUpLogger()
	r.Router()

}

func setUpLogger() {
	// Set up logger
	flag := log.Ldate
	datetime := time.Now().UTC().Format(YYYYMMDD+" "+HHMMSS12h) + ": "
	log.SetFlags(flag)
	log.SetPrefix(datetime)
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger := log.New(file, "", flag)
		logger.SetPrefix("FATAL: " + datetime)
		logger.Println(err)
	}
	defer file.Close()
	logger := log.New(file, "", flag)
	logger.SetPrefix("INFO : " + datetime)

	mw := io.MultiWriter(os.Stdout, file)

	logger.SetOutput(mw)
	logger.Println("Starting the company API")
}
