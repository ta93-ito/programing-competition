package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	x := 1

	result := 0

	for n >= x*2 {
		x *= 2
		result++
	}

	fmt.Println(result)
}
