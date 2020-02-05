package proxy

import (
	"fmt"
)

// CreateProfile will have the user define a reverse proxy
// profile that will be stored in the config directory.
func CreateProfile(proxy string) {
	var name string = proxy
	var port string

	fmt.Printf("Please enter a port number for %s to listen on:\n", name)
	fmt.Scanln(&port)

	emptyApps := make([]App, 0)

	newProfile := Profile{
		Name:       name,
		Port:       port,
		AppServers: emptyApps,
	}

	StoreProfile(&newProfile)
	fmt.Printf("Reverse Proxy profile *%s* has been saved!\n", name)
}
