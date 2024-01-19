package main

import (
	"fmt"
	"math"
)

// Under the hood, interface values can be thought of as
// a tuple of a value and a concrete type
type I interface {
	M()
}

// T is a concrete type that implements the I interface
type T struct {
	S string
}

// M implements the interface method for T
func (t *T) M() {
	fmt.Println(t.S)
}

// F is another concrete type that implements I
type F float64

// M implements the interface method for F
func (f F) M() {
	fmt.Println(f)
}

// describe prints details about an I interface value
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	// i is declared as an interface type I
	var i I

	// i is initialized with a *T value that implements I
	i = &T{"Hello"}

	// describe prints the value and concrete type
	describe(i)
	// (&{Hello}, *main.T)

	// The interface calls the implementation's M()
	i.M()
	// Hello

	// i is reassigned to an F value implementing I
	i = F(math.Pi)

	// describe prints the new value and concrete type
	describe(i)
	// (3.141592653589793, main.F)

	// The interface now calls F's implementation of M()
	i.M()
	// 3.141592653589793
}

// My Interface Values Example:
