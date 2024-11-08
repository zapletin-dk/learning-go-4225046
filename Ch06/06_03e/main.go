package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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

	tours := toursFromJson(content)
	for _, tour := range tours {
		fmt.Println(tour.Name)
	}

}

func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		if err != nil {
			panic(err)
		}
		tours = append(tours, tour)
	}
	return tours
}

type Tour struct {
	Name, Price string
}
