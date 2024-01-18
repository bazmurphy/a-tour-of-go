package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Functions that take a value argument must take a value of that specific type
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// while methods with value receivers take either a value or a pointer as the receiver when they are called
func (v Vertex) AbsMethod() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// v is a value, an instance of the Vertex struct
	v := Vertex{3, 4}

	// p is a pointer, to an instance of the Vertex struct
	p := &Vertex{4, 3}

	fmt.Println(v.AbsMethod())
	// 5

	// method's value receivers can take either a value or a pointer
	fmt.Println(p.AbsMethod())
	// 5

	// the function takes the value
	fmt.Println(AbsFunc(v))
	// 5

	// fmt.Println(AbsFunc(&v))
	// cannot use &v (value of type *Vertex) as Vertex value in argument to AbsFunccompilerIncompatibleAssign

	// the function takes a pointer
	fmt.Println(AbsFunc(*p))
	// 5
}
