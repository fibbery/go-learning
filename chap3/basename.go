package main

import (
	"fmt"
	"strings"
)

func basename(path string) string {
	path = path[strings.LastIndex(path, "/") + 1 :]
	if dot := strings.LastIndex(path,"."); dot >= 0 {
		path = path[:dot]
	}
	return path
}

func main(){
	fmt.Println(basename("a/b/c.go.go"))
}
