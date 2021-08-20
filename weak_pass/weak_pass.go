package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a string
	fmt.Scan(&a)
	b := strings.Split(a, "")
	var c [4]int
	for i := 0; i < 4; i++ {
		c[i], _ = strconv.Atoi(b[i])
	}
	same := true
	step := true
	for i := 0; i < 3; i++ {
		if c[i] != c[i+1] {
			same = false
		}
		if c[i+1]%10 != (c[i]+1)%10 {
			step = false
		}
	}
	if same || step {
		fmt.Println("Weak")
	} else {
		fmt.Println("Strong")
	}
}
