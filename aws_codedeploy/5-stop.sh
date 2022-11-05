#!/usr/bin/env bash -x

# This file is intended for AWS CodeDeploy service
# Referenced by "appspec.yml"
[[ -z `pgrep flight-finder` ]] || pkill flight-finder