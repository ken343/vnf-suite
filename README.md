McSecurity
==========

McSecurity is a suite of programs designed to emulate network functions virtually.
These inclue:

1. mcProxy - A CLI tool that can create a reverse proxy that will route requests to a list of application servers.
2. mcFilter = A virtual firewall.

mcProxy
-------

Usage:

```bash
mcproxy <SUBCOMMAND> [ARGUEMENTS]
```

Sub-Cmd | Explanation
--------|------------
add | adds a reverse proxy profile
remove | removes a reverse proxy profile
instance | list out all available reverse proxy profiles

### add

mproxy add _proxy-name_ _proxy-port_

#### Requirements

* Business:
Develop your own networking programs to watch, modify, and route network traffic to your HTTP server apps. Functions may be implemented as one application binary, or as several binaries.
* Solution:
    1. Functional
    2. Quality
        - [ ] Documentation
        - [ ] Agile Project Management
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
