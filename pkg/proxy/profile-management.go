package proxy

import (
	"encoding/json"
	"log"
	"os"
)

// LoadProfile looks up a proxy profile stored in config
// and returns the contents as a *Profile type.
func LoadProfile(profileName string) *Profile {
	file := profileName
	f, err := os.OpenFile(file, os.O_RDWR, 0777)
	if err != nil {
		log.Fatalf("Error Opening File: %v\n", err)
	}
	defer f.Close()

	newProfile := &Profile{}
	jsonerr := json.NewDecoder(f).Decode(newProfile)
	if jsonerr != nil {
		log.Fatalf("Error Decoding File: %v\n", err)
	}
	return newProfile
}

// StoreProfile takes an existing proxy profile and stores
// its settings within the config directory.
func StoreProfile(profile *Profile) {
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
