package main

import "fmt"

func main() {
	for {
		go fmt.Print(1)
		fmt.Print(0)
	}
}
