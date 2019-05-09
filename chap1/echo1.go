package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep = "", " "
	for i := 1; i < len(os.Args); i++{
		s = s + os.Args[i] + sep
	}
	fmt.Println(s)
}

