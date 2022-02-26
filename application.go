// This entry point is intended for AWS BeanStalk with Golang
package main

import (
	"flag"
	"log"

	"github.com/mateuszmidor/FlightFinder/cmd/finder_web/app"
)

func main() {
	port := flag.String("port", "5000", "-port=5000")
	flights_data_dir := flag.String("flights_data", "./assets", "-flights_data=./assets")
	web_data_dir := flag.String("web_data", "./web", "-web_data=./web")
	flag.Parse()

	log.SetPrefix("[APP] ")
	app.Run(*port, *flights_data_dir, *web_data_dir)
}
