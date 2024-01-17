package main

import "fmt"

func main() {
	// declares a variable `a` as an array of two strings
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// declares a variable `primes` as an array of 6 integers
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// Hello World
// [Hello World]
// [2 3 5 7 11 13]
