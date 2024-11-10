package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {

	// define client using default params
	client := http.Client{}
	
	// Create request for the client to perform
	req, err := http.NewRequest("GET", url, nil)
	checkError(err)

	// Set the User-Agent field in request header.
	req.Header.Set("User-Agent", "")

	// Perform the request and capture the response.
	resp, err := client.Do(req)
	checkError(err)
	defer resp.Body.Close()
	fmt.Printf("Response type: %T\n", resp)

	bytes, err := io.ReadAll(resp.Body)
	checkError(err)

	content := string(bytes)
	fmt.Print(content)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
