package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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
		Port: "7777",
		AppServers: []proxy.App{
			proxy.NewApp("localhost", "8081", "/"),
			proxy.NewApp("localhost", "8081", "/english"),
			proxy.NewApp("localhost", "8082", "/spanish"),
			proxy.NewApp("localhost", "8083", "/russian"),
		},
	}
)

func main() {

	log.SetFlags(log.Llongfile)

	// Test store and load functions.
	storeProfile(testProfile)
	mcProfile := loadProfile("./config/ken.json")

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

func loadProfile(profileName string) *proxy.Profile {
	file := profileName
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatalf("Error Opening File: %v\n", err)
	}
	defer f.Close()

	newProfile := &proxy.Profile{}
	jsonerr := json.NewDecoder(f).Decode(newProfile)
	if jsonerr != nil {
		log.Fatalf("Error Decoding File: %v\n", err)
	}
	return newProfile
}

func storeProfile(profile *proxy.Profile) {
	file := "./config/" + profile.Name + ".json"
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatalf("Error opening file to be written: %v\n", err)
	}
	defer f.Close()

	jsonerr := json.NewEncoder(f).Encode(profile)
	if jsonerr != nil {
		log.Fatalf("Error encoding file %v.\n", jsonerr)
	}
}
