package main

import "fmt"

func main() {
	m := make(map[string]int)

	// reassign a value to a specific key
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	// reassign a value to a specific key
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// delete a value of a specific key
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// check if a key exists
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

// The value: 42
// The value: 48
// The value: 0
// The value: 0 Present? false
