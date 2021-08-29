package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// 枝刈りで極力計算減らしてみたけど無理だった

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)

	n := getI()
	m := getI()
	a := getInts(n)
	b := getInts(m)

	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })

	minv := 1000000000

	for i, j := 0, 0; i < n && j < m; {
		minv = abs(a[i] - b[j])
		if a[i] > b[j] {
			j++
		} else {
			i++
		}
	}
	fmt.Println(minv)
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

func chmin(a *int, b int) bool {
	if *a < b {
		return false
	}
	*a = b
	return true
}
