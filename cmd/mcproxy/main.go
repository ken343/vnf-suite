package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Need to figure out what constants I need.
const (
	RPORT = ":8080" // RPORT for Reverse Proxy port to listen on. Replace with proxy type.
)

// Variable need to create a client where the reverse proxy can forward requests.
var (
	mcClient = &http.Client{}
)

func main() {
	fmt.Println("Howdy?")

	myProxyMux := http.NewServeMux()

	myProxyMux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {

		targetQuery := "?" + req.URL.RawQuery
		targetURL := "http://" + "localhost" + ":8081" + targetQuery
		fmt.Printf("Sanity Check targetURL == %s\n", targetURL)
		mcReq, err := http.NewRequest(req.Method, targetURL, nil)
		if err != nil {
			log.Fatalf("Endpoint /english new request not generated: %v\n", err)
		}
		defer req.Body.Close()

		resp, err := mcClient.Do(mcReq) //GET
		if err != nil {
			log.Fatalf("Endpoint /english new response not generated from application server: %v\n", err)
		}
		defer resp.Body.Close()

		buffer := make([]byte, 64)
		for {
			n, err := resp.Body.Read(buffer)
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatalf("Reading Error: %v", err)
			}
			fmt.Printf("Read %d bytes - buffer == %v", n, buffer[:n])

			fmt.Fprintf(rw, string(buffer))
		}

		fmt.Fprint(rw, "\n")
		fmt.Printf("Response Body is : %s\n", buffer)

		fmt.Printf("The port is ->%s\n", req.RemoteAddr)
	})

	myProxyMux.HandleFunc("/english", func(rw http.ResponseWriter, req *http.Request) {

		targetQuery := "?" + req.URL.RawQuery
		targetURL := "http://" + "localhost" + ":8081" + targetQuery
		fmt.Printf("Sanity Check targetURL == %s\n", targetURL)
		mcReq, err := http.NewRequest(req.Method, targetURL, nil)
		if err != nil {
			log.Fatalf("Endpoint /english new request not generated: %v\n", err)
		}
		defer req.Body.Close()

		resp, err := mcClient.Do(mcReq) //GET
		if err != nil {
			log.Fatalf("Endpoint /english new response not generated from application server: %v\n", err)
		}
		defer resp.Body.Close()

		buffer := make([]byte, 64)
		for {
			n, err := resp.Body.Read(buffer)
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatalf("Reading Error: %v", err)
			}
			fmt.Printf("Read %d bytes - buffer == %v", n, buffer[:n])

			fmt.Fprintf(rw, string(buffer))
		}

		fmt.Fprint(rw, "\n")
		fmt.Printf("Response Body is : %s\n", buffer)

		fmt.Printf("The port is ->%s\n", req.RemoteAddr)

	})

	myProxyMux.HandleFunc("/spanish", func(rw http.ResponseWriter, req *http.Request) {

		targetQuery := "?" + req.URL.RawQuery
		targetURL := "http://" + "localhost" + ":8082" + targetQuery
		fmt.Printf("Sanity Check targetURL == %s\n", targetURL)
		mcReq, err := http.NewRequest(req.Method, targetURL, nil)
		if err != nil {
			log.Fatalf("Endpoint /spanish new request not generated: %v\n", err)
		}
		defer req.Body.Close()

		resp, err := mcClient.Do(mcReq) //GET
		if err != nil {
			log.Fatalf("Endpoint /spanish new response not generated from application server: %v\n", err)
		}
		defer resp.Body.Close()

		buffer := make([]byte, 64)
		for {
			n, err := resp.Body.Read(buffer)
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatalf("Reading Error: %v", err)
			}
			fmt.Printf("Read %d bytes - buffer == %v", n, buffer[:n])

			fmt.Fprintf(rw, string(buffer))
		}

		fmt.Fprint(rw, "\n")
		fmt.Printf("Response Body is : %s\n", buffer)

		fmt.Printf("The port is ->%s\n", req.RemoteAddr)

	})

	myProxyMux.HandleFunc("/russian", func(rw http.ResponseWriter, req *http.Request) {

		targetQuery := "?" + req.URL.RawQuery
		targetURL := "http://" + "localhost" + ":8083" + targetQuery
		fmt.Printf("Sanity Check targetURL == %s\n", targetURL)
		mcReq, err := http.NewRequest(req.Method, targetURL, nil)
		if err != nil {
			log.Fatalf("Endpoint /russian new request not generated: %v\n", err)
		}
		defer req.Body.Close()

		resp, err := mcClient.Do(mcReq) //GET
		if err != nil {
			log.Fatalf("Endpoint /russian new response not generated from application server: %v\n", err)
		}
		defer resp.Body.Close()

		buffer := make([]byte, 64)
		for {
			n, err := resp.Body.Read(buffer)
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatalf("Reading Error: %v", err)
			}
			fmt.Printf("Read %d bytes - buffer == %v", n, buffer[:n])

			fmt.Fprintf(rw, string(buffer))
		}

		fmt.Fprint(rw, "\n")
		fmt.Printf("Response Body is : %s\n", buffer)

		fmt.Printf("The port is ->%s\n", req.RemoteAddr)

	})

	// myProxyServer will utilize the port indicated by the proxy profile.
	myProxyServer := &http.Server{
		Addr:         RPORT,
		Handler:      myProxyMux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	myProxyServer.ListenAndServe()

}
