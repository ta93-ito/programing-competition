package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	var result []int

	for i := 1; i <= num; i++ {
		if num%i == 0 {
			result = append(result, i)
		}
	}
	fmt.Println(result)
}
