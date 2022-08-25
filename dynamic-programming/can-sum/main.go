package main

import (
	"fmt"
	"log"
	"strconv"
)

func canSum(targetSum int, nums []int, combo string, memo map[int]bool) bool {
	// fmt.Printf("targetSum : %d\n", targetSum)
	if targetSum == 0 {
		log.Printf("%s : Found Combo\n", combo)
		return true
	}
	if targetSum < 0 {
		log.Printf("%s : Invalid Combo\n", combo)
		return false
	}
	if _, ok := memo[targetSum]; ok {
		log.Printf("%s : Memo Found\n", combo)
		return false
	}

	for _, num := range nums {
		rem := targetSum - num
		newCombo := combo + " : " + strconv.Itoa(num)
		if canSum(rem, nums, newCombo, memo) {
			return true
		} else {
			memo[rem] = false
		}
	}
	return false
}

func main() {
	memo := make(map[int]bool)
	fmt.Println(canSum(300, []int{7, 14}, " ", memo))

}
