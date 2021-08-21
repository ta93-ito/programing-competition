package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Scan(&n)

	s := make([]int, n)
	t := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(r, &s[i])
	}

	for i := 0; i < n; i++ {
		fmt.Fscan(r, &t[i])
	}

	for i := 0; i < 2*n; i++ {
		from := i % n
		next := (i + 1) % n
		t[next] = min(t[from]+s[from], t[next])

	}
	for i := 0; i < n; i++ {
		fmt.Println(t[i])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
