package main

import (
	"fmt"
	"strconv"
)

var (
	set = make(map[string]bool)
)

func main() {
	set["jiangnenghua"] = true
	set["jiangnenghua"] = true
	if set["jiangnenghua"] {
		fmt.Println("set has element jiangnenghua")
		fmt.Println("set size is " + strconv.Itoa(len(set)))
	}

	for j := 0; j < 64; j++{
		fmt.Println(uint(j))
	}
}
