package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// カッコの整合性チェック。
// カッコの組み合わせは、"()", "{}", "[]"
// 標準入力で与えられたカッコ列が成立するかどうか判定。
// "({)[]}"は成立する。

var sc = bufio.NewScanner(os.Stdin)
var runes = make([]string, 0)

func main() {
	sc.Split(bufio.ScanRunes)

	N := getI()

	// sc.Scanの挙動調べる
	// 0個めの要素に空文字が入る
	for i := 0; i < N; i++ {
		sc.Scan()
		new := sc.Text()
		fmt.Println(i, new)
		if !judgeBracket(new) {
			runes = append(runes, new)
		}
	}
	if len(runes) == 0 {
		fmt.Println(true)
	}
	fmt.Println(false)
}

func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

// judge and unset brackets if found reverse bracket
func judgeBracket(b string) bool {
	var reverseB string
	switch b {
	case ")":
		reverseB = "("
	case "}":
		reverseB = "}"
	case "]":
		reverseB = "]"
	}
	bI, isFound := findReverseBracket(reverseB)
	if isFound {
		remove(runes, bI)
		return true
	}
	return false
}

func findReverseBracket(b string) (int, bool) {
	for i := 0; i < len(runes); i++ {
		if runes[i] == b {
			return i, true
		}
	}
	return 0, false
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
