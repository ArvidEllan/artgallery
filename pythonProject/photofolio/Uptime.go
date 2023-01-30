package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	serverURL = "http://isalrightworldofart.com"
	timeout   = 5 * time.Second
)

func main() {
	for {
		// Set up an HTTP client with a timeout
		client := http.Client{
			Timeout: timeout,
		}

		// Make an HTTP GET request to the server
		response, err := client.Get(serverURL)
		if err != nil {
	
		// The request failed, so the server is down
			fmt.Println("Server is down:", err)
		} else {
			// The request succeeded, so the server is up
			fmt.Println("Server is up:", response.Status)
		}

		// Wait before checking again
		time.Sleep(30 * time.Second)
	}
}