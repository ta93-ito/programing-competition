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
	a := strings.Split(getS(), ".")
	x, _ := strconv.Atoi(a[0])
	y, _ := strconv.Atoi(a[1])
	if y <= 2 {
		fmt.Println(fmt.Sprintf("%d-", x))
	} else if y <= 6 {
		fmt.Println(x)
	} else {
		fmt.Println(fmt.Sprintf("%d+", x))
	}
}

func getS() string {
	sc.Scan()
	return sc.Text()
}
