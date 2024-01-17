package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	// A struct literal denotes a newly allocated struct value by listing the values of its fields
	// You can list just a subset of fields by using the Name: syntax (And the order of named fields is irrelevant)
	v1 = Vertex{1, 2}  // has type Vertex
	p  = &Vertex{1, 2} // has type *Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
)

func main() {
	fmt.Println(v1, p, v2, v3)
}

// {1 2} &{1 2} {1 0} {0 0}
