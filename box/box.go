package main

import (
	"fmt"
	"strconv"
)

func main() {

	var input [3]string
	fmt.Scan(&input[0], &input[1], &input[2])

	var n, a, b int
	n, _ = strconv.Atoi(input[0])
	a, _ = strconv.Atoi(input[1])
	b, _ = strconv.Atoi(input[2])

	fmt.Println(n - a + b)
}
