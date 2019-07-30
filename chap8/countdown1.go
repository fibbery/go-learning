package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)
	for i := 10; i > 0; i-- {
		fmt.Println(i)
		<-tick
	}
	fmt.Println("launch")
}
