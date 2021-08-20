package main

import "fmt"

func main() {
	var x, y, a, b int
	fmt.Scan(&x, &y, &a, &b)

	ex := 0

	for {
		if x*a > x+b || x*a >= y {
			break
		}
		x *= a
		ex++
	}
	ex += (y - 1 - x) / b
	fmt.Println(ex)
}
