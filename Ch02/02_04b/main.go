package main

import "fmt"

func main() {
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 +i3
	fmt.Println("Integer sum:", intSum)

	f1, f2, f3 := 0.3, 0.3, 0.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum:", floatSum)
}
