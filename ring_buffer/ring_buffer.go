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
	buf := make([]int, 5)
	n := getI()

	for i := 0; i < n; i++ {
		bi := i % 5
		buf[bi] = getI()
		fmt.Println(buf[bi])
	}
	fmt.Println(buf)
}

func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
