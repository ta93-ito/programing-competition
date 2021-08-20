package main

import "fmt"

func main() {
	var gold, silver int
	fmt.Scan(&gold, &silver)
	if gold == 0 {
		fmt.Println("Silver")
	} else if silver != 0 {
		fmt.Println("Alloy")
	} else {
		fmt.Println("Gold")
	}
}
