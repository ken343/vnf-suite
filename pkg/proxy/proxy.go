package proxy

// Proxy will hold the port of the current reverse proxy server
// as well as load in new array of application servers.
type Proxy struct {
	Port       string
	AppServers []App
}

// App represents the configuration data for mcProxy will be
// redirecting traffic to.
type App struct {
	Host     string
	Port     string
	Endpoint string
}
