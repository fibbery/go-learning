package main

import (
	"fmt"
	"io"
)

func main() {
	var d io.Writer
	f(d)
}

func f(out io.Writer) {
	fmt.Printf("%T\n", out)
	if out != nil {
		out.Write([]byte("hello\n"))
	}
}
