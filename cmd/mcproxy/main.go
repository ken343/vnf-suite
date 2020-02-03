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
// This variable is declared globally because documentation says that http.Client
// is safe for concurrent use and uses caching. Thus all routes will use a single
// mcClient.
var (
	mcClient = &http.Client{}
)

func main() {
	fmt.Printf("Howdy, mcProxy is listening on port %s\n...", RPORT)

	myProxyMux := http.NewServeMux()

	defaultRoute := NewRoute("localhost", "8081", "/", mcClient)
	myProxyMux.Handle("/", defaultRoute)

	englishRoute := NewRoute("localhost", "8081", "/english", mcClient)
	myProxyMux.Handle("/english", englishRoute)

	spanishRoute := NewRoute("localhost", "8082", "/spanish", mcClient)
	myProxyMux.Handle("/spanish", spanishRoute)

	russianRoute := NewRoute("localhost", "8083", "/russian", mcClient)
	myProxyMux.Handle("/russian", russianRoute)

	// myProxyServer will utilize the port indicated by the proxy profile.
	myProxyServer := &http.Server{
		Addr:         RPORT,
		Handler:      myProxyMux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	myProxyServer.ListenAndServe()

}
