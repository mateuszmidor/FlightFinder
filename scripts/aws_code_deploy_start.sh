#!/usr/bin/env bash

# This file is intended for AWS CodeDeploy service
# Referenced by "appspec.yml"
cd /FlightFinder/
go build -o flight-finder cmd/finder_web/main.go

# Run as daemon
setsid ./flight-finder -flights_data=./assets -web_data=./web -port=80 >/FlightFinder/flight-finder.logs 2>&1 < /FlightFinder/flight-finder.logs &
