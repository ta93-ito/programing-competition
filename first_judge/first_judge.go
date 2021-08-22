package main

import "fmt"

const a = "Hello,World!"

func main() {
	var s string
	fmt.Scan(&s)
	if s == a {
		fmt.Println("AC")
	} else {
		fmt.Println("WA")
	}
}
