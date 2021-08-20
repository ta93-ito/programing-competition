package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([]int, n)
	b := make([]int, m)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	for i := 0; i < m; i++ {
		fmt.Scan(&b[i])
	}

	sort.Ints(a)
	sort.Ints(b)

	result := int(math.Pow10(9)) - 1

	for i, j := 0, 0; i < n && j < m; {
		if abs(a[i], b[j]) < result {
			result = abs(a[i], b[j])
		}
		if a[i] > b[j] {
			j++
		} else if a[i] <= b[j] {
			i++
		}
	}
	fmt.Println(result)
}

func abs(a int, b int) int {
	if a-b < 0 {
		return (a - b) * -1
	} else {
		return a - b
	}
}
