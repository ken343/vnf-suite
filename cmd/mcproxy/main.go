package main

import (
	"fmt"
	"log"
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

		targetQuery := "?" + req.URL.RawQuery
		targetURL := "http://" + "localhost" + ":8083" + targetQuery
		fmt.Printf("Sanity Check targetURL == %s\n", targetURL)
		mcReq, err := http.NewRequest(req.Method, targetURL, req.Body)
		if err != nil {
			log.Fatalf("Endpoint /russian new request not generated: %v\n", err)
		}
		defer req.Body.Close()

		resp, err := mcClient.Do(mcReq)
		if err != nil {
			log.Fatalf("Endpoint /russian new response not generated from application server: %v\n", err)
		}
		defer resp.Body.Close()

		fmt.Fprintf(rw, "This is the final return %v", resp.Body)

		fmt.Printf("The port is ->%s\n", req.RemoteAddr)

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
