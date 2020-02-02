package main

import (
	"fmt"
	"net/http"
	"time"
)

// Need to figure out what constants I need.
const (
	RPORT = ":8080" // RPORT for Reverse Proxy port to listen on. Replace with proxy type.
)

// Variable need to create a client where the reverse proxy can forward requests.
var (
	mcClient = &http.Client{}
)

func main() {
	fmt.Println("Howdy?")

	myProxyMux := http.NewServeMux()
	// myProxyMux.HandleFunc("/")
	// myProxyMux.HandleFunc("/english")
	// myProxyMux.HandleFunc("/spanish")
	myProxyMux.HandleFunc("/russian", func(rw http.ResponseWriter, req *http.Request) {
		mcReq := req
		fmt.Printf("The port is ->%s\n", mcReq.URL.Port())
	})

	// myProxyServer will utilize the port indicated by the proxy profile.
	myProxyServer := &http.Server{
		Addr:         RPORT,
		Handler:      myProxyMux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	myProxyServer.ListenAndServe()
}
