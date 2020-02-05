package proxy

import (
	"fmt"
	"log"
	"os"
)

// UnMount removes an application server from the selected
// proxy profile.
func UnMount(proxy string) {
	mcProxy := LoadProfile("./config/" + proxy + ".json")

	// Print out selection and prompt for input.
	fmt.Printf("Pofile Selected: %s\n", proxy)
	fmt.Println("=========================================")
	for i := 0; i < len(mcProxy.AppServers); i++ {
		fmt.Printf("[%d] App Server: %v\n", i, mcProxy.AppServers[i])
	}
	fmt.Print("\n=========================================\n")
	fmt.Println("[Enter an integer to select app server for unmounting.]")
	fmt.Print("Enter: ")
	var x int
	fmt.Scanln(&x)
	fmt.Print("\n")

	// Remove app from slice.
	mcProxy.AppServers = RemoveIndex(mcProxy.AppServers, x)

	// Delete file so that a new one can be re-written.
	err := os.Remove("./config/" + proxy + ".json")
	if err != nil {
		log.Fatalf("Error in File replacement: %v", err)
	}

	// Store new and improved Profile.
	StoreProfile(mcProxy)
}

// RemoveIndex deletes App from a slice.
func RemoveIndex(s []App, index int) []App {
	return append(s[:index], s[index+1:]...)
}
