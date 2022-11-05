#!/usr/bin/env bash -x

# This file is intended for AWS CodeDeploy service
# Referenced by "appspec.yml"
cd /FlightFinder/
go build -o flight-finder cmd/finder_web/main.go
