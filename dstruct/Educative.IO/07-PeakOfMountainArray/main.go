package main

import "fmt"

func main() {
	inp := []int{50, 60, 70, 80, 90, 10, 20, 30, 40}
	res := peakMountainOfArray(inp)
	fmt.Printf("index : %d, value: %d \n", res, inp[res])
}

func peakMountainOfArray(arr []int) int {
	left, right, targetIndex := 0, len(arr)-1, -1

	for left < right {

	}
	return targetIndex
}
