package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 入力 N:要素数 Array: 配列

var sc = bufio.NewScanner(os.Stdin)

const Sentinel = 1000000001
const Max = 500000

func main() {
	sc.Split(bufio.ScanWords)
	N := getI()
	target := getInts(N)

	mergeSort(target, N, 0, N)
	fmt.Println(target)
}

func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getI()
	}
	return ret
}

func mergeSort(A []int, n, left, right int) {
	if left+1 < right {
		mid := (left + right) / 2
		mergeSort(A, n, left, mid)
		mergeSort(A, n, mid, right)
		merge(A, n, left, mid, right)
	}
}

func merge(A []int, n, left, mid, right int) {
	n1 := mid - left
	n2 := right - mid

	L := make([]int, 5)
	R := make([]int, 5)

	for i := 0; i < n1; i++ {
		L[i] = A[left+i]
	}
	for i := 0; i < n2; i++ {
		R[i] = A[mid+i]
	}
	L[n1], R[n2] = Sentinel, Sentinel
	i, j := 0, 0
	for k := left; k < right; k++ {
		if L[i] <= R[j] {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
	}
}
