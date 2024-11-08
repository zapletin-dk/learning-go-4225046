package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {

	// define client using default params
	client := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}
	// Create request for the client to perform
	req, err := http.NewRequest("GET", url, nil)
	// Set the User-Agent field in request header.
	req.Header.Set("User-Agent", "")

	if err != nil {
		panic(err)
	}
	// Perform the request and capture the response.
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response type: %T\n", resp)
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	content := string(bytes)
	fmt.Print(content)
}
