package main

import (
	"fmt"
	"math"
)

func main() {
	// In Go, a name is exported if it begins with a capital letter
	// Pi is an exported name, exported from the math package
	// When importing a package you can only refer to its exported names
	fmt.Println(math.Pi)
}
// 3.141592653589793