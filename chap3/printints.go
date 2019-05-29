package main

import (
	"bytes"
	"fmt"
)

func intsToString(values []int) string{
	var buf bytes.Buffer
	buf.WriteString("[")
	for index,value := range values {
		if index > 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf,"%d", value)
	}
	buf.WriteString("]")
	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3,}))
	fmt.Printf("%b\n",4)
}
