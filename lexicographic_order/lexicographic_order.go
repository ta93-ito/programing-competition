package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	s := getS()
	t := getS()

	if SIsShort(s, t) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func getS() string {
	sc.Scan()
	return sc.Text()
}

func SIsShort(s, t string) bool {
	for i := 0; ; i++ {
		if i == len(s) || i == len(t) {
			return len(s) < len(t)
		}

		if s[i] == t[i] {
			continue
		}
		return s[i] < t[i]
	}
}
