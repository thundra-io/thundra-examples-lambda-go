#!/bin/bash
GOOS=linux GOARCH=amd64 go build -o helloworld helloworld.go
sls deploy