package main

import (
	"fmt"
	"math"
)

// convert the type with Type(value) T(v)
func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	// var f float64 = math.Sqrt(x*x + y*y)
	// cannot use x * x + y * y (value of type int) as float64 value in argument to math.SqrtcompilerIncompatibleAssign
	var z uint = uint(f)
	// var z uint = f
	// cannot use f (variable of type float64) as uint value in variable declarationcompilerIncompatibleAssign
	fmt.Println(x, y, z)
}
// 3 4 5