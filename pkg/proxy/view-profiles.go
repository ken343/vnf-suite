package proxy

import (
	"fmt"
	"os/exec"
	"strings"
)

// ViewProfiles shows what reverse proxies are available
// and what application servers are currently available for routing.
func ViewProfiles() {

	// Use linux find command to acquire list of json profiles
	// in /config
	find := exec.Command("find", "config/")

	found, err := find.Output()
	if err != nil {
		fmt.Printf("Error, could not get file list: %v", err)
	}

	// Take binary output of find cmd, convert to string and
	// split at new characters.
	separated := strings.Split(string(found), "\n")
	var altered []string = make([]string, 0)

	// Add ./ to each file to create relative path.
	for i := 0; i < len(separated); i++ {
		altered = append(altered, "./"+separated[i])
	}

	// Prune empty config and blank './' from list so that whole slice can
	// be used to load profiles.
	n := len(altered)
	pruned := altered[1 : n-1]

	// Create slice of loaded profiles.
	sProfiles := make([]*Profile, 0)
	for i := 0; i < len(pruned); i++ {
		profile := LoadProfile(pruned[i])
		sProfiles = append(sProfiles, profile)
	}

	// Print out Profiles for viewing. Supreme success.
	fmt.Println("=========================================")
	for i, v := range sProfiles {
		fmt.Printf("[%d] %s listening on Port: %s\n", i, v.Name, v.Port)
		for _, a := range sProfiles[i].AppServers {
			fmt.Printf("\t| %v|\n", a)
		}
	}
	fmt.Print("=========================================\n")
}
