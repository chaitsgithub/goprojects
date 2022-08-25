package main

import "fmt"

func factorial(i int) int {
	if i == 0 {
		return 1
	}
	return i * factorial(i-1)
}

func main() {
	x := 5
	fmt.Printf("Factorial of %v is %v", x, factorial(x))
}
