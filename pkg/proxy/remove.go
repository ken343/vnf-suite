package proxy

import (
	"fmt"
	"log"
	"os/exec"
)

// Remove will delete the selected profile from the
// config directory.
func Remove(proxy string) {
	profileFile := "./config/" + proxy + ".json"
	deleteProfile := exec.Command("rm", profileFile)

	myBytes, err := deleteProfile.Output()
	if err != nil {
		log.Fatalf("Error deleting file %v  -> %v", myBytes, err)
	}

	fmt.Printf("%s has been deleted...\n", proxy)
}
