#!/usr/bin/env bash -x

# This file is intended for AWS CodeDeploy service
# Referenced by "appspec.yml"
curl --fail localhost:80 > index.html