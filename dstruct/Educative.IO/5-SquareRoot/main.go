package main

import "fmt"

func main() {
	fmt.Println(squareRoot(4))
	fmt.Println(squareRoot(9))
	fmt.Println(squareRoot(10))
	fmt.Println(squareRoot(17))
}

func squareRoot(target int) int {
	left, right, result := 0, target, 0
	for left <= right {
		mid := (left + right) / 2
		square := mid * mid
		if square > target {
			right = mid - 1
		} else {
			left = mid + 1
			result = mid
		}
	}
	return result
}
