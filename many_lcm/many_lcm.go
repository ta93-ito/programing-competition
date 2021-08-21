package main

import "fmt"

// N個の数の最小公倍数を求める

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < n; {
		if len(arr) == 1 {
			break
		}
		tmp := lcm_of_double(arr[i], arr[i+1])
		arr[i] = tmp
		arr = remove(arr, i+1)
	}
	fmt.Println(arr[0])
}

func lcm_of_double(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for {
		if b == 0 {
			break
		}
		tmp := b
		b = a % b
		a = tmp
	}
	return a
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
