package main

import "fmt"

type Vertex struct {
	X, Y float64
}

// methods with pointer receivers take either a value or a pointer as the receiver when they are called
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// functions with a pointer argument must take a pointer
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	// v is a value, an instance of the Vertex struct
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)
	fmt.Println(v)
	// {60 80}

	// p is a pointer, to an instance of the Vertex struct
	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)
	fmt.Println(p)
	// &{96 72}
}
