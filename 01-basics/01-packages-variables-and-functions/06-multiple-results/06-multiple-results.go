package main

import "fmt"

// swap function takes two strings as parameters and returns them in reverse order.
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// Call the swap function with "hello" and "world" as arguments, and receive the results in variables a and b.
	// := is shorthand for declaring and assigning to variables
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
// world hello