#!/usr/bin/env bash 

# This file is intended for AWS CodeDeploy service
# Referenced by "appspec.yml"
set -x
curl --fail localhost:80 > index.html