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
	s := getS()
	if s[n-1] == 'o' {
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

func getS() string {
	sc.Scan()
	return sc.Text()
}
