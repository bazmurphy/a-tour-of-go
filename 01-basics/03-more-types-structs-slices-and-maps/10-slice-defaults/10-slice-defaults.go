package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	// (!) mutating the array each time

	// s[1:4] creates a new slice from index 1 to (4-1) = 3
	// so it includes elements at indices 1, 2, and 3.
	s = s[1:4]
	fmt.Println(s) // Output: [3 5 7]

	// s[:2] creates a new slice from the beginning of the slice to index (2-1) = 1
	// so it includes elements at index 0 and 1.
	s = s[:2]
	fmt.Println(s) // Output: [3 5]

	// s[1:] creates a new slice from index 1 to the end of the slice.
	s = s[1:]
	fmt.Println(s) // Output: [5]

}

// [3 5 7]
// [3 5]
// [5]
