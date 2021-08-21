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
	snukes := make([]int, n)
	for i := 0; i < len(snukes); i++ {
		snukes[i] = t[i]
	}
	for i := 0; i < len(t); i++ {
		for j := 0; j < len(s); j++ {
			if snukes[(j+1)%n] > snukes[j]+s[j] {
				snukes[(j+1)%n] = snukes[j] + s[j]
			}
		}
	}
	for i := 0; i < len(snukes); i++ {
		fmt.Println(snukes[i])
	}
}
