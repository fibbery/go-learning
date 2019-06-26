package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
	"unicode/utf8"
)

func equal(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}

	for k, av := range a {
		if bv, ok := b[k]; !ok || bv != av {
			return false
		}
	}

	return true
}

func main() {
	defer costtime("main")()
	a, b := make(map[string]int), make(map[string]int)
	a["bob"] = 1
	b["bob"] = 1
	fmt.Println(equal(a, b))
	x, y := map[string]int{"A": 0}, map[string]int{"B": 42}
	fmt.Println(x["B"] == y["B"])
	fmt.Println([]rune("hello中"))
	fmt.Println(utf8.RuneCountInString("中国"))
	var hello func()

	hello = func() {
		fmt.Println("hello,world")
		hello()
	}

	double(3)

}

func costtime(methodname string) func() {
	start := time.Now()
	log.Printf("enter the method %s at %v\n", methodname, start)
	return func() {
		log.Printf("exit the method %s at %v, method costtime is %d \n", methodname, time.Now(), time.Since(start).Nanoseconds()/1000)
	}
}

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	runtime.Stack()
	return x * x
}
