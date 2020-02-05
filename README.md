McSecurity
==========

McSecurity is a suite of programs designed to emulate network functions virtually.
These inclue:

1. mcProxy - A CLI tool that can create a reverse proxy that will route requests to a list of application servers.
2. mcFilter - A virtual firewall.

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
unmount | removes a busienss server that proxy has mounted
run | runs instance of selected profile


### add
```bash
$ mproxy add _proxy-name_ _proxy-port_
```
#### Requirements

* Business:
Develop your own networking programs to watch, modify, and route network traffic to your HTTP server apps. Functions may be implemented as one application binary, or as several binaries.
* Solution:
    1. Functional
    2. Quality
        - [ ] Documentation
        - [X] Agile Project Management
        - [ ] Unit Testing
        - [ ] Logs & Metrics
        - [ ] Environment Configuration
        - [ ] Secuirty
        - [ ] Build & Deploy Scripts
        - [ ] Containerization
    3. Performance
* Stakeholder & Legal:
    * 10 Minute Demonstration
    * Presentation Slides
