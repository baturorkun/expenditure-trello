#! /usr/bin/env bash
########################### GITLAB CREDENTIALS ###############################

set -e

export GOARCH=amd64
export GOOS=linux

#export GOPATH=/builder
#mkdir -p $PWD/build/{frontend,scripts,docs}

go build -v -o ./expenditure expenditure
