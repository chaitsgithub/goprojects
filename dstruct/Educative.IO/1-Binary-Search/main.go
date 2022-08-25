package main

import (
	"log"
)

func main() {
	i1 := []int{1, 3, 5, 7, 9, 11, 13, 15}
	log.Println(binarySearch(i1, 5))
	log.Println(binarySearch(i1, 4))
	i2 := []int{}
	log.Println(binarySearch(i2, 5))
}

func binarySearch(arr []int, target int) bool {
	left, right := 0, len(arr)-1

	for right >= left {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return true
		} else if arr[mid] > target {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		}
	}
	return false
}
