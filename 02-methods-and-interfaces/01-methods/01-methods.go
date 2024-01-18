package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Go does not have classes. However, you can define methods on types
// A method is a function with a special receiver argument
// The receiver appears in its own argument list between the func keyword and the method name
func (v Vertex) Abs() float64 {
	// 	^ receiver argument
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

// 5

// My Methods Example:

// type Person struct {
// 	name string
// 	age  int
// }

// func (p Person) sayHello() {
// 	fmt.Printf("Hello I am %v and i am %d\n", p.name, p.age)
// }

// func (p Person) sayGoodbye() {
// 	fmt.Printf("Goodbye!\n")
// }

// func main() {
// 	var baz Person
// 	baz.name = "Baz"
// 	baz.age = 99
// 	baz.sayHello()
// 	baz.sayGoodbye()
// }

// Hello I am Baz and i am 99
// Goodbye!
