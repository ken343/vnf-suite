package proxy

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Default multiplexor used by the proxy package.

var myProxyMux = http.NewServeMux()

// Profile will hold the port of the current reverse proxy server
// as well as load in new array of application servers.
type Profile struct {
	Name       string
	Port       string
	AppServers []App
}

// App Contains the data support data that creates a handler
// that will dynamically add several application server routes
// based on a profile. My swap this with proxy type in /pkg/proxy/
type App struct {
	AppHost  string
	AppPort  string
	Endpoint string
}

// NewApp constructs a new App type.
func NewApp(host string, port string, endpoint string) App {
	newApp := App{
		AppHost:  host,
		AppPort:  port,
		Endpoint: endpoint,
	}

	return newApp

}

// String implements the "Stringer" interface.
func (r App) String() string {
	return fmt.Sprintf("%s --> %s:%s", r.Endpoint, r.AppHost, r.AppPort)
}

// ServeHTTP implements the Handler interface.
func (r App) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	// Create URL that targets the appropriate application server.
	targetQuery := "?" + req.URL.RawQuery
	targetURL := "http://" + r.AppHost + ":" + r.AppPort + r.Endpoint + targetQuery

	fmt.Println("====================================================================")
	fmt.Printf("Routing to IP == %s\n", r.AppHost)
	fmt.Printf("Routing to Port == %s\n", r.AppPort)
	fmt.Printf("Endpoint mcProxy received == %s\n", r.Endpoint)
	fmt.Printf("Sanity Check targetURL == %s\n", "http://"+r.AppHost+":"+r.AppPort+targetQuery)
	fmt.Println("--------------------------------------------------------------------")

	// Generate a new request that will be sent to the target application
	// server.
	mcReq, err := http.NewRequest(req.Method, targetURL, req.Body)
	if err != nil {
		log.Fatalf("Endpoint /%s new request not generated: %v\n", r.Endpoint, err)
	}
	defer req.Body.Close()

	// Send request to target application server.
	resp, err := http.DefaultClient.Do(mcReq) //GET
	if err != nil {
		log.Fatalf("Endpoint /%s new response not generated from application server: %v\n", r.Endpoint, err)
	}
	defer resp.Body.Close()

	rw.Header().Set("Content-Type", "text/plain")
	// Read contents of response body into the new response writer
	// that will be sent to the original client.
	buffer := make([]byte, 64)
	rw.Write([]byte("\n"))
	for {
		n, err := resp.Body.Read(buffer)
		if err != nil {
			if err == io.EOF {
				rw.Write(bytes.Trim(buffer, "\x00"))

				fmt.Printf("Read %d bytes into buffer == %v\n", n, buffer[:n])
				break
			}

			log.Fatalf(err.Error())

		}
		fmt.Printf("Read %d bytes -> buffer == %v\n", n, buffer[:n])

		rw.Write(bytes.Trim(buffer, "\x00"))
	}

	// Add a newline for readability and print some sanity checks.
	// Sanity checks may be removed at a later date.
	rw.Write([]byte("\n"))
	fmt.Printf("Response Body is : %s\n", buffer)

	fmt.Printf("The origin host:port is == %s\n", req.RemoteAddr)
	fmt.Println("====================================================================")
}

// AddServer uses NewApp behind the scenes to add another application
// server to the Profiles array.
func (p *Profile) AddServer(host string, port string, endpoint string) {
	newApp := NewApp(host, port, endpoint)
	p.AppServers = append(p.AppServers, newApp)
}
