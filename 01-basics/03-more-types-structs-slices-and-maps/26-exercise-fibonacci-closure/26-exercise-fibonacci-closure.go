package main

import "fmt"

// fibonacci is a function that returns a function that returns an int.
func fibonacci() func() int {
	// initialize current and next
	current := 0
	next := 1

	// return a function closure
	return func() int {
		// store current in num variable
		num := current

		// swap values
		current = next
		next = num + next

		// return num
		return num
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
