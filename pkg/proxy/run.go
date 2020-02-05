package proxy

import (
	"fmt"
	"net/http"
	"time"
)

// Run starts running a Server based on the selected profile.
func Run(proxy string) {
	mcProfile := LoadProfile("./config/" + proxy + ".json")

	fmt.Printf("Howdy, mcProxy is listening on port %s...\n", ":"+mcProfile.Port)
	// Remember that proxy.Profile.Appservers[i] implement the "Handler" interface.
	// Hence they are being attached to the multiplexor.
	for i := 0; i < len(mcProfile.AppServers); i++ {
		myProxyMux.Handle(mcProfile.AppServers[i].Endpoint, mcProfile.AppServers[i])
	}

	// myProxyServer will utilize the port indicated by the proxy profile.
	myProxyServer := &http.Server{
		Addr:         ":" + mcProfile.Port,
		Handler:      myProxyMux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	myProxyServer.ListenAndServeTLS("cert.pem", "key.pem")

}
