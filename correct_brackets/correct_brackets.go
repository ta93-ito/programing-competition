package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

// judge whether given bracket is correct
// correct cases: "{}", "[[[({)}]]]"
// incorrect cases: "{{", "((([[]))]]"

func main() {
	sc.Split(bufio.ScanWords)
	getI()
	bracketList := strings.Split(getS(), "")
	bracketStack := make([]string, 0)
	for _, b := range bracketList {
		i, stackable := judgeStackable(bracketStack, b)
		if stackable {
			bracketStack = append(bracketStack, b)
		} else {
			bracketStack = remove(bracketStack, i)
		}
	}
	if len(bracketStack) == 0 {
		fmt.Println("Given bracket is correct!")
	} else {
		fmt.Println("Given bracket is incorrect!")
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

// return stackable and pair's index. return -1 index if stackable.
// false -> remove pair without stack true -> stack
func judgeStackable(arr []string, elem string) (int, bool) {
	// resolve pair
	var target string
	switch elem {
	case ")":
		target = "("
	case "}":
		target = "{"
	case "]":
		target = "["
	default:
		return -1, true
	}

	i, isFound := find(arr, target)
	if isFound {
		return i, false
	}
	return -1, true
}

func find(arr []string, elem string) (int, bool) {
	// for i, e := range arr {
	// 	if e == elem {
	// 		return i, true
	// 	}
	// }
	// return -1, false
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == elem {
			return i, true
		}
	}
	return -1, false
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
