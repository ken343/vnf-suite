package proxy

import "fmt"

// AddApplication takes the selected proxy file and "mounts"
// the business server to its routing.
func AddApplication(proxy string) {
	var host string
	var port string
	var endpoint string
	var answer string = "n"

	selectedProxy := LoadProfile("./config/" + proxy + ".json")
	for {
		fmt.Println("Please enter the domain name of the application server you would like to add to profile:")
		fmt.Scanln(&host)
		fmt.Printf("Please enter the port that %s will be listening on:\n", host)
		fmt.Scanln(&port)
		fmt.Printf("Please enter proxy endpoint that will act as alias for %s:%s (include leading slash):\n", host, port)
		fmt.Scanln(&endpoint)

		newApp := NewApp(host, port, endpoint)
		selectedProxy.AppServers = append(selectedProxy.AppServers, newApp)
		fmt.Println("New application mount successful.")
		fmt.Println("Would you like to add another? [y/n]")
		fmt.Scanln(&answer)
		if answer == "n" {
			StoreProfile(selectedProxy)
			fmt.Println("Transaction complete.")
			return
		}
	}
}
