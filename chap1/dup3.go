package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, e := ioutil.ReadFile(filename)
		if e != nil {
			_, _ = fmt.Fprintf(os.Stderr, "open file error, %t\n", e)
			continue
		}
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			counts[line]++
		}
		for line,count := range counts {
			fmt.Println()
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}
