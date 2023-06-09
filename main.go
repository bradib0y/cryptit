package main

import (
	"fmt"

	"github.com/bradib0y/cryptit/encrypt"
)

func main() {
	fmt.Println(encrypt.Nimbus("Kodekloud"))
}

// go mod init: create a module in the current directory (go.mod file)
// go mod tidy: ensures go.mod file is matching the .go source code requirements
// go run: builds and runs a file, then cleans up artifacts (good for testing)
// go build: compiles the packages and creates an executable in current directory
// go install: compiles and creates executable to $GOPATH/bin
// go env: show Go environment variables
// go get: get specific package versions, and add them to go.mod, and download to module cache
