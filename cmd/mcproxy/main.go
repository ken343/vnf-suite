package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ken343/vnf-suite/pkg/proxy"
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
	fmt.Printf("Howdy, mcProxy is listening on port %s...\n", RPORT)

	myProxyMux := http.NewServeMux()

	defaultApp := proxy.NewApp("localhost", "8081", "/", mcClient)
	myProxyMux.Handle("/", defaultApp)

	englishApp := proxy.NewApp("localhost", "8081", "/english", mcClient)
	myProxyMux.Handle("/english", englishApp)

	spanishApp := proxy.NewApp("localhost", "8082", "/spanish", mcClient)
	myProxyMux.Handle("/spanish", spanishApp)

	russianApp := proxy.NewApp("localhost", "8083", "/russian", mcClient)
	myProxyMux.Handle("/russian", russianApp)

	// myProxyServer will utilize the port indicated by the proxy profile.
	myProxyServer := &http.Server{
		Addr:         RPORT,
		Handler:      myProxyMux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	myProxyServer.ListenAndServe()

}
