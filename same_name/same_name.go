package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanLines)
	n := getI()
	ns := getStrings(n)
	if isDuplicate(ns) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getStrings(N int) []string {
	ret := make([]string, N)
	for i := 0; i < N; i++ {
		ret[i] = getS()
	}
	return ret
}

func getS() string {
	sc.Scan()
	return sc.Text()
}

func isDuplicate(l []string) bool {
	m := map[string]bool{}
	for _, e := range l {
		if m[e] == true {
			return true
		} else {
			m[e] = true
		}
	}
	return false
}
