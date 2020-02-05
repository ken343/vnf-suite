package main

import (
	"flag"
	"fmt"
	"log"
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
// is safe for concurrent use and uses caching.
var (
	myProxyMux  = http.NewServeMux()
	testProfile = &proxy.Profile{
		Name: "ken",
		Port: "8080",
		AppServers: []proxy.App{
			proxy.NewApp("localhost", "8081", "/"),
			proxy.NewApp("localhost", "8081", "/english"),
			proxy.NewApp("localhost", "8082", "/spanish"),
			proxy.NewApp("localhost", "8083", "/russian"),
		},
	}
)

func main() {
	// Set up error log to ouptut information
	log.SetFlags(log.Lshortfile)

	flag.Parse()
	fmt.Printf("flag.Arg(0) == %v\n", flag.Arg(0))
	switch flag.Arg(0) {
	case "create":
		proxy.CreateProfile()
	case "mount":
		proxy.AddApplication(flag.Arg(1))
	default:
		fmt.Println("No valid sub command selected.")
	}

	// Test store and load functions.
	proxy.StoreProfile(testProfile)
	mcProfile := proxy.LoadProfile("./config/ken.json")

	fmt.Printf("Howdy, mcProxy is listening on port %s...\n", ":"+mcProfile.Port)
	// Remember that proxy.Profile.Appservers[i] implement the "Handler" interface.
	// Hence they are being attachec to the multiplexor.
	myProxyMux.Handle(mcProfile.AppServers[0].Endpoint, mcProfile.AppServers[0])
	myProxyMux.Handle(mcProfile.AppServers[1].Endpoint, mcProfile.AppServers[1])
	myProxyMux.Handle(mcProfile.AppServers[2].Endpoint, mcProfile.AppServers[2])
	myProxyMux.Handle(mcProfile.AppServers[3].Endpoint, mcProfile.AppServers[3])

	// myProxyServer will utilize the port indicated by the proxy profile.
	myProxyServer := &http.Server{
		Addr:         ":" + mcProfile.Port,
		Handler:      myProxyMux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	myProxyServer.ListenAndServe()

}
