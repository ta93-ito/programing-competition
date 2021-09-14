package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	given_quiz := make([]string, 3)
	for i := 0; i < 3; i++ {
		given_quiz[i] = getS()
	}
	all_quiz := []string{"ABC", "ARC", "AGC", "AHC"}
	fmt.Println(unique(all_quiz, given_quiz)[0])
}

func unique(ss ...[]string) []string {
	m := map[string]int{}
	for _, s := range ss {
		for _, v := range s {
			m[v]++ // 出現回数をカウント
		}
	}
	res := []string{}
	for k, v := range m {
		if v == 1 {
			// 出現回数が１回のものだけを抽出
			res = append(res, k)
		}
	}
	return res
}

func getS() string {
	sc.Scan()
	return sc.Text()
}
