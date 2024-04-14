package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Define the target URL
	url := "https://www.bing.com/search?q=Quote%20of%20the%20day"

	// Create a new HTTP client
	client := &http.Client{}

	// Make an HTTP GET request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send the request and get the response
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close() // Ensure proper closing of the body

	// Check for successful response (status code 200)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:", resp.StatusCode)
		return
	}

	// Read the response body into a byte slice
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the first few bytes of the response to show content retrieval
	fmt.Println("Fetched content (first 500 bytes):")
	fmt.Println(string(body[:500])) // Limit output to 500 bytes
}
