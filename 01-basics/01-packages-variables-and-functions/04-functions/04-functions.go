package main

import "fmt"

// declare the type after the parameter name
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
// 55