package main

import (
	"bufio"
	"strings"
)

type WordCounter int

func (w *WordCounter) Write(p []byte) (n int, err error) {
	str := strings.NewReader(string(p))
	bs := bufio.NewScanner(str)
	bs.Split(bufio.ScanWords)
	sum := 0
	for bs.Scan(){
		sum++
	}
	*w = WordCounter(sum)
	return sum, nil
}



