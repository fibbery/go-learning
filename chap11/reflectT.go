package main

import (
	"fmt"
	"reflect"
)

func main() {
	var param = struct {
		width  int
		height int
	}{}
	fmt.Printf("%T\n", param)
	fmt.Println(reflect.TypeOf(param).Kind() == reflect.Struct)
	fmt.Println(reflect.ValueOf(param))

	var mod = 3
	reflect.ValueOf(&mod).Elem().Set(reflect.ValueOf(4))
	fmt.Println(mod)
}
