package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	maxp := -200000
	minv := a[0]
	for i := 1; i < n; i++ {
		maxp = max(maxp, a[i]-minv)
		minv = min(minv, a[i])
	}
	fmt.Println(maxp)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
