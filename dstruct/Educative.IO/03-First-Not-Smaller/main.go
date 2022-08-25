package main

import "log"

func main() {
	inp := []int{1, 3, 5, 7, 9, 11, 13, 15}
	log.Println(firstNotSmaller(inp, 5))
	log.Println(firstNotSmaller(inp, 6))
	log.Println(firstNotSmaller(inp, 0))
}

func firstNotSmaller(arr []int, target int) int {
	left, right := 0, len(arr)-1
	index := -1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] >= target {
			index = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return index
}
