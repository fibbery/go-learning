package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x interface{} = 3
	b := reflect.ValueOf(&x).Elem()
	fmt.Println(b.CanAddr())
}
