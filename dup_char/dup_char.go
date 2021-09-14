package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

// 隣接する文字が同じであればどちらも削除し、残った文字列を出力する。
// 入力 "aabccbdf" 出力 "df"

func main() {
	sc.Split(bufio.ScanWords)
	ans := removeDupChar(getS())
	fmt.Println(ans)
}

func getS() string {
	sc.Scan()
	return sc.Text()
}

func removeDupChar(s string) string {
	sa := strings.Split(s, "")
	st := make([]string, 1)
	st = append(st, sa[0])

	for i := 1; i < len(sa); i++ {
		if len(st) == 0 {
			st = append(st, sa[i])
			continue
		}
		// スタックの最後の要素とsa[i]が同じならスタックから取り除く
		if st[len(st)-1] == sa[i] {
			st = st[:len(st)-1]
		} else {
			st = append(st, sa[i])
		}
	}
	return strings.Join(st, "")
}
