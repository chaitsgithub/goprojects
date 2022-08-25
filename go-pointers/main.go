package main

import "log"

func main() {
	a := 2
	log.Printf("Value: %v, Address: %v\n", a, &a)

	square1(&a)
	log.Printf("Value: %v, Address: %v\n", a, &a)
}

func square1(a *int) {
	*a *= *a
	log.Printf("a: %v, *a: %v, &a: %v\n", a, *a, &a)
}
