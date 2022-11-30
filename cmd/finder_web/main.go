package main

import (
	"flag"
	"log"

	"github.com/mateuszmidor/FlightFinder/cmd/finder_web/app"
)

func main() {
	log.SetPrefix("[APP] ")

	port := flag.String("port", "8080", "-port=80")
	flights_data_dir := flag.String("flights_data", "./assets", "-flights_data=./assets")
	web_data_dir := flag.String("web_data", "./web", "-web_data=./web")
	aws_region := flag.String("aws_region", "us-east-1", "-aws-region=us-east-1")
	aws_xray := flag.Bool("aws_xray", false, "-aws_xray=false")
	redis_addr := flag.String("redis_addr", "localhost:6379", "-redis_addr=localhost:6379")
	redis_pass := flag.String("redis_pass", "CACHE", "-redis_pass=CACHE")
	flag.Parse()

	app.Run(*port, *flights_data_dir, *web_data_dir, app.MakeMetricsClient(*aws_region), app.MakeTracingClient(*aws_xray), app.MakeCacheClient(*redis_addr, *redis_pass))
}
