package main

import (
	"fmt"
)

func main() {
	inp := []int{50, 60, 70, 80, 90, 100, 10, 20, 30, 40}
	res := findMinRotated(inp)
	fmt.Printf("index : %d, value: %d \n", res, inp[res])

	inp = []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	res = findMinRotated(inp)
	fmt.Printf("index : %d, value: %d \n", res, inp[res])

	inp = []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	res = findMinRotated(inp)
	fmt.Printf("index : %d, value: %d \n", res, inp[res])

	inp = []int{}
	res = findMinRotated(inp)
	if res == -1 {
		fmt.Printf("Empty Array")
	} else {
		fmt.Printf("index : %d, value: %d \n", res, inp[res])
	}
}

func findMinRotated(arr []int) int {
	left, right, targetIndex := 0, len(arr)-1, -1

	for left < right {
		mid := (left + right) / 2
		if arr[mid] > arr[right] {
			left = mid + 1
		} else {
			right = mid
			targetIndex = mid
		}
	}

	return targetIndex
}
