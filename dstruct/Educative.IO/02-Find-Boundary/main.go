package main

import "log"

func main() {
	arr1 := []bool{false, false, true, true, true}
	arr2 := []bool{false, false, false, false, false}
	arr3 := []bool{true, true, true, true, true, true}
	arr4 := []bool{}
	log.Println(findBoundary(arr1))
	log.Println(findBoundary(arr2))
	log.Println(findBoundary(arr3))
	log.Println(findBoundary(arr4))
}

func findBoundary(arr []bool) int {
	boundaryIndex := -1
	left, right := 0, len(arr)-1

	for left < right {
		mid := (left + right) / 2
		if arr[mid] {
			boundaryIndex = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return boundaryIndex
}
