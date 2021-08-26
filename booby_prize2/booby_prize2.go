package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Player struct {
	number int
	score  int
}

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()

	a := make([]Player, n)
	for i := 0; i < n; i++ {
		a[i] = Player{
			score:  nextInt(),
			number: i + 1,
		}
	}

	max1, max2 := Player{score: 0, number: 0}, Player{score: 0, number: 0}

	for i := 0; i < n; i++ {
		if a[i].score > max1.score {
			max2 = max1
			max1 = a[i]
		} else if a[i].score > max2.score {
			max2 = a[i]
		}
	}
	fmt.Println(max2.number)
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func nextStr() string {
	sc.Scan()
	return sc.Text()
}
