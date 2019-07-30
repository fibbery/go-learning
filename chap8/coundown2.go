package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	tick := time.Tick(1 * time.Second)
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	for i := 10; i > 0; i-- {
		select {
		case <-tick:
			fmt.Println(i)
		case <-abort:
			fmt.Println("abort")
			return
		}
	}

	fmt.Println("launch!!!!")
}
