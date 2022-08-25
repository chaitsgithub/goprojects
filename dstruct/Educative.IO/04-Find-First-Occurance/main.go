package main

import "log"

func main() {
	inp := []int{1, 3, 3, 3, 3, 6, 10, 10, 10, 100}
	log.Println(findFirstOccurance(inp, 1))
	log.Println(findFirstOccurance(inp, 3))
	log.Println(findFirstOccurance(inp, 6))
	log.Println(findFirstOccurance(inp, 10))
	log.Println(findFirstOccurance(inp, 100))
}

func findFirstOccurance(arr []int, target int) int {
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
