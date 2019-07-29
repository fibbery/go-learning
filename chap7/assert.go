package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	_, ok := w.(*os.File)
	fmt.Println(ok)
	os.Open()
}
