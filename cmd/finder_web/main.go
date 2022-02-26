package main

import (
	"flag"
	"log"

	"github.com/mateuszmidor/FlightFinder/cmd/finder_web/app"
)

func main() {
	log.SetPrefix("[APP] ")
	port := flag.String("port", "8080", "-port=80")
	flights_data_dir := flag.String("flights_data", "./assets", "-data=./assets")
	web_data_dir := flag.String("web_data", "./web", "-web_data=./web_data")
	flag.Parse()

	app.Run(*port, *flights_data_dir, *web_data_dir)
}
