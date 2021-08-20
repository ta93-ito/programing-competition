package main

import (
	"fmt"
	"os"
)

func main() {
	var path string
	fmt.Scan(&path)

	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	data := make([]byte, 1024)
	_, err2 := f.Read(data)
	if err2 != nil {
		fmt.Println(err2)
	}

	defer f.Close()

	fmt.Println(string(data))
}
