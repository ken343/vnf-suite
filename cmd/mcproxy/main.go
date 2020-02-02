package main

import (
	"fmt"
	"net/http"
)

// Need to figure out what constants I need.
const (
	RPORT = ":8080" // RPORT for Reverse Proxy port to listen on. Replace with proxy type.
)

func main() {
	fmt.Println("Howdy?")

	http.ListenAndServe(RPORT, nil)
}
