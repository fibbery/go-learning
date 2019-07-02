package main

import "fmt"

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (n int, err error) {
	*b += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0
	fmt.Fprintf(&c,"hello, %s", "jiangnenghua")
	fmt.Println(c)
}


