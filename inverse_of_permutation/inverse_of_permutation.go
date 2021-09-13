package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	n := getI()
	a := getInts(n)

	b := make([]string, n+1)

	for i := 0; i < n; i++ {
		b[a[i]] = strconv.Itoa(i + 1)
	}
	fmt.Println(strings.Join(b[1:], " "))
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
