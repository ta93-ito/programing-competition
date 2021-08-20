package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n <= 125 {
		fmt.Println(4)
	} else if n <= 211 {
		fmt.Println(6)
	} else {
		fmt.Println(8)
	}
}
