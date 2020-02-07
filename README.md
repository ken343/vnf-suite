McSecurity
==========

McSecurity is a suite of programs designed to emulate network functions virtually.
These inclue:

1. mcProxy - A CLI tool that can create a reverse proxy that will route requests to a list of application servers.

mcProxy
-------

Usage:

```bash
$ mcproxy <SUBCOMMAND> [ARGUEMENTS]
```

Sub-Cmd | Explanation
--------|-----------
build | creates a reverse proxy profile that can be mounted to
mount | adds a business server that proxy should route to
unmount | removes a business server that proxy has mounted
run | runs instance of selected profile
profiles | list out mcProxy profiles
remove | removes a profile from memory

