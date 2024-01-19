package main

import (
	"fmt"
	"math"
)

// An interface type is defined as a set of method signatures
// A value of interface type can hold any value that implements those methods
type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex) and does NOT implement Abser.
	// Vertex (the value type) doesn't implement Abser because the Abs method is defined only on *Vertex (the pointer type).
	a = v

	fmt.Println(a.Abs())
}

// My Interfaces Example:

// type PersonStruct struct {
// 	name string
// 	age  int
// }

// type PersonInterface interface {
// 	sayHello()
// 	sayGoodbye()
// }

// // method with pointer receiver
// func (p *PersonStruct) sayHello() {
// 	fmt.Printf("Hello %v\n", p.name)
// }

// // method with pointer receiver
// func (p *PersonStruct) sayGoodbye() {
// 	fmt.Printf("Goodbye %v\n", p.name)
// }

// func main() {
// 	// create a struct
// 	ps := PersonStruct{name: "Baz", age: 99}

// 	// create an interface
// 	var pi PersonInterface

// 	// pointer the interface to the struct
// 	// interfaces reference values
// 	pi = &ps

// 	fmt.Println(ps)
// 	// {Baz 99}

// 	fmt.Println(pi)
// 	// &{Baz 99}

// 	// now we can call the interface methods on the struct (because the interface proxies the calls)
// 	ps.sayHello()
// 	ps.sayGoodbye()
// }
