package main

import (
	"fmt"
	"strconv"
	"strings"
)

func comma(number string) string {
	n := len(number)
	if n <= 3 {
		return number
	}
	return strings.Join([]string{comma(number[:n-3]),number[n-3:]},",")
}

func main() {
	fmt.Println(comma(strconv.Itoa(1111111111)))
	fmt.Println(strconv.Atoi("1111111111"))

}
