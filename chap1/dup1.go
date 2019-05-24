package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}

	for text, count := range counts {
		fmt.Printf("%d\t%s\n", count, text)
	}
}
