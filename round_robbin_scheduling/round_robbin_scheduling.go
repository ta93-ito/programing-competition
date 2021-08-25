package main

import (
	"fmt"
)

type Process struct {
	name string
	time int
}

func main() {
	// 入力 :プロセス数N, クオンタイムQ, ["name(i) time(i)"]
	var n, q int
	fmt.Scan(&n, &q)

	pa := make([]Process, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&pa[i].name)
		fmt.Scan(&pa[i].time)
	}

	var result []string
	t := 0
	for len(pa) > 0 {
		if pa[0].time-q <= 0 {
			t += pa[0].time
			result = append(result, fmt.Sprintf("%s %d", pa[0].name, t))
			pa = pa[1:]
		} else {
			t += q
			pa[0].time = pa[0].time - q
			v := pa[0]
			pa = append(pa[1:], v)
		}
	}
	for i := 0; i < n; i++ {
		fmt.Println(result[i])
	}
}
