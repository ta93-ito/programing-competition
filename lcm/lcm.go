package main

import "fmt"

func main() {

	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(a * b / gcd(a, b))
}

func gcd(a, b int) int {
	for {
		if b == 0 {
			break
		}
		tmp := b
		b = a % b
		a = tmp
	}
	return a
}
