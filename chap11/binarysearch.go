package main

import "fmt"

func binarySearch(data []int, leftIndex, rightIndex, searchData int) {
	midIndex := (rightIndex + leftIndex) / 2
	if searchData > data[midIndex] {
		binarySearch(data, midIndex+1, rightIndex, searchData)
	} else if searchData < data[midIndex] {
		binarySearch(data, leftIndex, midIndex-1, searchData)
	} else {
		fmt.Printf("找到对应数据，位于index : %d\n", midIndex)
		return
	}
	if leftIndex >= rightIndex {
		fmt.Println("未能寻找到对应数据")
		return
	}
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7,}
	binarySearch(data, 0, len(data)-1, 7)
}
