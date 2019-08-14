package main

import (
	"fmt"
)

func main() {
	var data struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}
	fmt.Print(data)
}
