package main

import (
	"fmt"
)

func howSum(targetSum int, nums []int, memo map[int]bool) (bool, []int) {
	// fmt.Printf("targetSum : %d\n", targetSum)
	if targetSum == 0 {
		return true, nil
	}
	if targetSum < 0 {
		return false, nil
	}
	if _, ok := memo[targetSum]; ok {
		return false, nil
	}

	for _, num := range nums {
		rem := targetSum - num
		res, comboNums := howSum(rem, nums, memo)
		if res {
			comboNums = append(comboNums, num)
			return true, comboNums
		} else {
			memo[rem] = false
		}
	}
	return false, nil
}

func main() {
	memo := make(map[int]bool)
	fmt.Println(howSum(300, []int{7, 14}, memo))
	fmt.Println(howSum(7, []int{2, 3}, memo))

}
