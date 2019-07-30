package main

import (
	"flag"
	"fmt"
	"strings"
)

type sliceValue []string

func (d *sliceValue) String() string {
	return strings.Join(*d, ",")
}

func (d *sliceValue) Set(args string) error {
	*d = strings.Split(args, ",")
	return nil
}

// to instantiation
func newSliceVar(origin []string, point *[]string) *sliceValue {
	*point = origin
	return (*sliceValue)(point)
}

func main() {
	var directories sliceValue
	flag.Var(&directories, "d", "directory")
	flag.Parse()
	fmt.Println(directories)
}
