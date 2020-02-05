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

	flag.Parse()
	fmt.Printf("flag.Arg(0) == %v\n", flag.Arg(0))
	switch flag.Arg(0) {
	case "build":
		proxy.CreateProfile()
	case "mount":
		proxy.Mount(flag.Arg(1))
	case "run":
		proxy.Run(flag.Arg(1))
	default:
		fmt.Println("No valid sub command selected.")
	}
}
