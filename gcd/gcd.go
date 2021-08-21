package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	for {
		if b == 0 {
			fmt.Println(a)
			break
		}
		tmp := b
		b = a % b
		a = tmp
	}
}
