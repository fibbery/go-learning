package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filenames := os.Args[1:]
	if len(filenames) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, filename := range filenames {
			file, e := os.Open(filename)
			if e != nil {
				fmt.Fprint(os.Stderr, "open file error, %v\n", e)
				continue
			}
			countLines(file, counts)
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

func countLines(file *os.File, counts map[string]int) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		counts[scanner.Text()]++
		if counts[scanner.Text()] > 1 {
			fmt.Printf("file %s has duplicate line, line is %s\n", file.Name(), scanner.Text())
		}
	}

}
