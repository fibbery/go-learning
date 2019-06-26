package main

import "fmt"

func throwPanic() {
	defer func(){
		if p := recover(); p != nil {
			fmt.Println("has occur error")
		}
	}()
	panic("this is error")
}

func main() {
	throwPanic()
}
