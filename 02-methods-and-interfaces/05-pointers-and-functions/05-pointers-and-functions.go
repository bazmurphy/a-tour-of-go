package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// this is no longer a method, it has no receiver. it is just a function
func ScaleByValue(v Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// this is no longer a method, it has no receiver. it is just a function
// the first parameter is a pointer to a Vertex
func ScaleByPointer(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	ScaleByValue(v, 10)
	fmt.Println(Abs(v))
	// 5

	v = Vertex{3, 4}
	ScaleByPointer(&v, 10) // get the memory address of v
	fmt.Println(Abs(v))
	// 50
}
