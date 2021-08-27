package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	_ = nextInt()
	_ = nextInt()
	n := nextInt()

	ha := make([]int, n)
	wa := make([]int, n)

	for i := 0; i < n; i++ {
		ha[i] = nextInt()
		wa[i] = nextInt()
	}

	cpha := make([]int, n)
	copy(cpha, ha)
	cpwa := make([]int, n)
	copy(cpwa, wa)

	sort.Slice(cpha, func(i, j int) bool { return cpha[i] < cpha[j] })
	sort.Slice(cpwa, func(i, j int) bool { return cpwa[i] < cpwa[j] })

	mapha := make(map[int]int)
	mapwa := make(map[int]int)

	ba := 0
	cnt := 0
	for i := 0; i < n; i++ {
		a := cpha[i]
		if ba == a {
			cnt++
			continue
		}
		mapha[a] = i + 1 - cnt
		ba = a
	}

	ba = 0
	cnt = 0
	for i := 0; i < n; i++ {
		a := cpwa[i]
		if ba == a {
			cnt++
			continue
		}
		mapwa[a] = i + 1 - cnt
		ba = a
	}

	for i := 0; i < n; i++ {
		h := ha[i]
		w := wa[i]
		fmt.Println(mapha[h], mapwa[w])
	}
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
