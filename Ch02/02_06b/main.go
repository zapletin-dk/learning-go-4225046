package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("time ", t)

	n := time.Now()
	fmt.Println("time now ", n)

	fmt.Printf(n.Format(time.ANSIC) + "\n")

	tommorow := n.AddDate(0, 0, 1)
	fmt.Printf(tommorow.Format(time.ANSIC) + "\n")

	format := "Mon 2006-02-01 03:02:01"
	fmt.Printf(tommorow.Format(format) + "\n")
}
