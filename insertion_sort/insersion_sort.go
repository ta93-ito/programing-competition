package main

import (
	"fmt"
)

func main() {
	arr := getArray()
	for i := 0; i < len(arr); i++ {
		v := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > v {
			arr[j+1] = arr[j]
			fmt.Println(arr)
			j--
		}
		fmt.Println(j)
		arr[j+1] = v
	}
	fmt.Println(arr)
}

func getArray() []int {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	return arr
}
