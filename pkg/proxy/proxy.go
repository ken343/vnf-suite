package proxy

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Profile will hold the port of the current reverse proxy server
// as well as load in new array of application servers.
type Profile struct {
	Port       string
	AppServers []App
}

// App Contains the data support data that creates a handler
// that will dynamically add several application server routes
// based on a profile. My swap this with proxy type in /pkg/proxy/
type App struct {
	appHost     string
	appPort     string
	endpoint    string
	routeClient *http.Client
}

// NewApp constructs a new App type.
func NewApp(host string, port string, endpoint string, client *http.Client) App {
	newApp := App{
		appHost:     host,
		appPort:     port,
		endpoint:    endpoint,
		routeClient: client,
	}

	return newApp

}

func (r App) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	// Create URL that targets the appropriate application server.
	targetQuery := "?" + req.URL.RawQuery
	fmt.Printf("r.endpoint == %s\n", r.endpoint)
	targetURL := "http://" + r.appHost + r.appPort + r.endpoint + targetQuery
	fmt.Printf("Sanity Check targetURL == %s\n", targetURL)

	// Generate a new request that will be sent to the target application
	// server.
	mcReq, err := http.NewRequest(req.Method, targetURL, req.Body)
	if err != nil {
		log.Fatalf("Endpoint /%s new request not generated: %v\n", r.endpoint, err)
	}
	defer req.Body.Close()

	// Send request to target application server.
	resp, err := r.routeClient.Do(mcReq) //GET
	if err != nil {
		log.Fatalf("Endpoint /%s new response not generated from application server: %v\n", r.endpoint, err)
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

	fmt.Printf("The sending port is == %s\n", req.RemoteAddr)

}
