#!/bin/sh

# comment out the following steps for the first build
# go get -u github.com/golang/dep/cmd/dep
# dep ensure
# go get -u github.com/gobuffalo/packr/...

GOOS=windows GOARCH=386 packr build -o generator.exe cmd/main.go
GOOS=windows GOARCH=amd64 packr build -o generator_64.exe cmd/main.go
packr build -o generator cmd/main.go
