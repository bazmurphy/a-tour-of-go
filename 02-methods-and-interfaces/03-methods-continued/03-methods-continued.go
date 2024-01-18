package main

import (
	"fmt"
	"math"
)

// You can declare a method on non-struct types, too
// In this example we see a numeric type `MyFloat`
type MyFloat float64

// with an `Abs` method
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
