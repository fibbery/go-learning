package main

import (
	"fmt"
	"net/url"
)

type Point struct {
	x, y int
}

func (p *Point) runByRef() {
	fmt.Printf("hello, x = %d, y=%d\n", p.x, p.y)
}

func (p Point) runByDirect() {
	fmt.Printf("hello, x = %d, y=%d\n", p.x, p.y)
}

func main() {
	p := Point{1, 2}
	paddress := &p
	paddress.runByDirect()
	paddress.runByRef()
	p.runByDirect()
	p.runByRef()
	values := url.Values(nil)
	fmt.Println(values["Hello"])
	runByRef := (*Point).runByRef
	runByDirect := Point.runByDirect
	fmt.Printf("%T\n%T\n", runByRef, runByDirect)

	var names= make(map[string]bool)
	names["jiangnenghua"]= true
	fmt.Println(names["lixin"])
}
