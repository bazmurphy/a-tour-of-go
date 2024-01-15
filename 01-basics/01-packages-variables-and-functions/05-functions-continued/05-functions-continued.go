package main

import "fmt"

// if consecutive parameters share the same type, you can add the type to the last
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
// 55