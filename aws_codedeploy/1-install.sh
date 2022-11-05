#!/usr/bin/env bash -x

# This file is intended for AWS CodeDeploy service
# Referenced by "appspec.yml"
yum update -y
yum install -y golang
mkdir -p /FlightFinder