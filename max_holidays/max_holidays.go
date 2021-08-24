package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, h int
	fmt.Scan(&n, &h)

	list := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&list[i])
	}
	fmt.Println(n, h, len(list))
	fi := Index(list, "work")
	max := 0

	if fi != -1 {
		for i := fi; i < n; i++ {

			l := make([]string, n)
			copy(l, list)
			r := h

			for j := i; j < n; j++ {
				if l[j] == "work" {
					l[j] = "off"
					r--
				}
				if r == 0 {
					break
				}
			}
			m := maxList(strings.Split(strings.Join(l, ""), "work")) / 3
			if m > max {
				max = m
			}
		}
		fmt.Println(max)
	} else {
		fmt.Println(n)
	}
}

func Index(l []string, q string) int {
	a := -1

	for i, v := range l {
		if v == q {
			a = i
			break
		}
	}

	return a
}

func maxList(l []string) int {
	max := 0
	for i := 0; i < len(l); i++ {
		if len(l[i]) > max {
			max = len(l[i])
		}
	}
	return max
}
