package proxy

import "fmt"

// UnMount removes an application server from the selected
// proxy profile.
func UnMount(proxy string) {
	mcProxy := LoadProfile("./config/" + proxy + ".json")

	fmt.Printf("Pofile Selected: %s\n", proxy)
	fmt.Println("=========================================")
	for i := 0; i < len(mcProxy.AppServers); i++ {
		fmt.Printf("[%d] App Server: %v\n", i, mcProxy.AppServers[i])
	}
	fmt.Print("\n=========================================\n")
	fmt.Println("[Enter an integer to select app server to unmount from profile.]")
	fmt.Print("Enter: ")
	var x int
	fmt.Scanln(&x)
	fmt.Print("\n")

	mcProxy.AppServers = RemoveIndex(mcProxy.AppServers, x)
	StoreProfile(mcProxy)

}

// RemoveIndex deletes App from a slice.
func RemoveIndex(s []App, index int) []App {
	return append(s[:index], s[index+1:]...)
}
