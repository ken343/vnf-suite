package main

import (
	"github.com/ken343/vnf-suite/pkg/proxy"
)

// testProfile is a backup profile that will be used in case of errors with
// mount sub-command.
var (
	testProfile = &proxy.Profile{
		Name: "backup",
		Port: "8080",
		AppServers: []proxy.App{
			proxy.NewApp("localhost", "8081", "/"),
			proxy.NewApp("localhost", "8081", "/english"),
			proxy.NewApp("localhost", "8082", "/spanish"),
			proxy.NewApp("localhost", "8083", "/russian"),
		},
	}
)

func init() {
	proxy.StoreProfile(testProfile)
}
