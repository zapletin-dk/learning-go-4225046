package main

import (
	"fmt"
	"time"
)

func say(message string, sleepDuration int) {
	time.Sleep(1 * time.Second)
	fmt.Println(message)
}

func main() {
	go say("Hello from the Goroutine", 1)
	println("Hello from Main")
	time.Sleep(2 * time.Second)
	println("All done")
}
