package main

import "fmt"

// If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

// interface defines a single method
type I interface {
	M()
}

// A "concrete" type in Go refers to a real, instantiated type like a struct, int, string etc. as opposed to an interface type
type T struct {
	S string
}

// T implements M() method required by interface I
func (t *T) M() {
	// In some languages this would trigger a null pointer exception,
	// but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method M in this example.)
	// Note that an interface value that holds a nil concrete value is itself non-nil.
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	// create an interface value i, initialized to nil
	var i I

	// create a pointer to T, initialized to nil
	var t *T

	// Assign the nil *T value to the interface
	i = t

	// i contains a nil *T concrete value
	describe(i)
	// (<nil>, *main.T)

	// Calling a method on a nil interface calls the method with a nil receiver
	i.M()
	// <nil>

	// Initialize i to a non-nil *T
	i = &T{"hello"}

	// Now describe prints a non-nil value
	describe(i)
	// (&{hello}, *main.T)

	// Calling the method now dispatches to the non-nil *T
	i.M()
	// hello
}
