package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// N
// A1 ... AN

type Player struct {
	number int
	score  int
}

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]Player, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i].score)
		a[i].number = i + 1
	}

	sort.Slice(a, func(i, j int) bool { return a[i].score < a[j].score })
	fmt.Println(a[n-2].number)
}
