package main

import "fmt"

func main() {
	// 出力 : 要素数N, 配列A
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	for i := 0; i < n; i++ {
		mini := i
		for j := i; j < n; j++ {
			if a[j] < a[mini] {
				mini = j
			}
		}
		a[i], a[mini] = a[mini], a[i]
	}
	fmt.Println(a)
}
