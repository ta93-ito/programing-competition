package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)

	n := getI()

	steps := getInts(n)

	dp := make([]int, n)
	dp[0] = 0

	for i := 1; i < n; i++ {
		if i == 1 {
			dp[i] = dp[i-1] + abs(steps[i-1]-steps[i])
		} else {
			dp[i] = min(dp[i-1]+abs(steps[i-1]-steps[i]),
				dp[i-2]+abs(steps[i-2]-steps[i]))
		}
	}
	fmt.Println(dp[n-1])
}

func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getI()
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
