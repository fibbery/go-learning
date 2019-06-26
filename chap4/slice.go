package main

import (
	"fmt"
)

func reverse(data *[5]int) {
	for i, j := 0, len(*data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func clearSame(data []int) []int {
	for i := 0; i < len(data)- 1; i++ {
		if data[i] == data[i+1] {
			copy(data[i:], data[i+2:])
			data = data[:len(data)-2]
			i = i - 2
		}
	}
	return data
}

func main() {
	test := [...]int{1, 2, 3, 4, 5, 5, 6, 7, 7, 9, 7,}
	fmt.Print(clearSame(test[:]))
}
