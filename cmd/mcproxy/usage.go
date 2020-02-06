package main

const usage = `
mcproxy is a tool for creating and running reverse proxies.

Usage:

        mcproxy <command> [arguments]

The commands are:

		build		compile packages and dependencies
		run 		compile and run Go program
		profiles	list out reverse proxy profiles
		remove		remove selected profile from system
		mount		add a routing to an application server
		unmount		remove a routing to an application server
`
