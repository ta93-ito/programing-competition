package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	flag := true

	for i := 0; flag; i++ {
		flag = false
		for j := n - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				tmp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = tmp
				flag = true
			}
		}
	}
	fmt.Println(arr)
}
