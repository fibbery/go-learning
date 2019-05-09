package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, v := range os.Args[1:] {
		s = s + v + sep
	}
	fmt.Println(s)
}