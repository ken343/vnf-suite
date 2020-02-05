package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ken343/vnf-suite/pkg/proxy"
)

func main() {
	// Set up error log to ouptut information
	log.SetFlags(log.Lshortfile)

	// Parse Command Line arguements and use them to find
	// appropriate sub-command for mcProxy to run.
	flag.Parse()
	switch flag.Arg(0) {
	case "build":
		proxy.CreateProfile(flag.Arg(1))
	case "mount":
		proxy.Mount(flag.Arg(1))
	case "run":
		proxy.Run(flag.Arg(1))
	case "unmount":
		proxy.UnMount(flag.Arg(1))
	default:
		fmt.Println("No valid sub command selected.")
	}
}
