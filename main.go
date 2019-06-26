package main

import (
	"bytes"
	"fmt"
)

func main() {
	var hash bytes.Buffer
	hash.WriteByte('c')
	hash.WriteByte('h')
	var bits = []int{16, 8, 4, 2, 1,}
	fmt.Println(bits[0])
}
