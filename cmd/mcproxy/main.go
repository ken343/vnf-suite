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

	RouteApp("/", "localhost", "8081", mcClient, myProxyMux)
	RouteApp("/english", "localhost", "8081", mcClient, myProxyMux)
	RouteApp("/spanish", "localhost", "8082", mcClient, myProxyMux)
	RouteApp("/russian", "localhost", "8083", mcClient, myProxyMux)

	// myProxyServer will utilize the port indicated by the proxy profile.
	myProxyServer := &http.Server{
		Addr:         RPORT,
		Handler:      myProxyMux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	myProxyServer.ListenAndServe()

}

// RouteApp takes the input data of the application server
// and hooks it up to the proxy multiplexer. The RouteApp function
// also requires an http.Client to make requests from.
func RouteApp(endpoint string, ip string, port string, client *http.Client, mux *http.ServeMux) {
	newApp := proxy.NewApp(ip, ":"+port, endpoint, client)
	mux.Handle(endpoint, newApp)
}
