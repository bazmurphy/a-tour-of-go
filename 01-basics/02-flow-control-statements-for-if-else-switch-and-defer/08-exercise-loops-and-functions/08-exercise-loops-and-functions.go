package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	// initialise z as a float 1.0
	z:= float64(1)

	// loop 10 times
	for i := 0; i < 10; i++ {
		// use Newton's method
		z -= (z * z - x) / ( 2 * z)
		// print a string template, pass the values at the end, and define their format inside the string template
		// %d is integer
		// %v is default format of the value
		fmt.Printf("own Sqrt method, iteration %d: z = %v\n", i+1, z)
	}

	return z
}

func main() {
	// compare the two methods accuracy
	fmt.Println(Sqrt(2))
	fmt.Println("math.Sqrt method, z = ", math.Sqrt(2))
}
