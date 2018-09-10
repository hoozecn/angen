#!/bin/sh

# comment out the following steps for the first build
# go get -u github.com/golang/dep/cmd/dep
# dep ensure

GOOS=windows GOARCH=386 go build -o angen.exe angen/main.go
GOOS=windows GOARCH=amd64 go build -o angen_64.exe angen/main.go
GOOS=darwin  GOARCH=amd64 go build -o angen_max angen/main.go
