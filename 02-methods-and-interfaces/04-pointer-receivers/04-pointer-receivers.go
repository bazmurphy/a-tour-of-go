package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// func (v *Vertex) Scale(f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

func (v Vertex) ScaleWithValueReceiver(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// You can declare methods with pointer receivers
// Methods with pointer receivers can modify the value to which the receiver points (as Scale does here)
// Since methods often need to modify their receiver, pointer receivers are more common than value receivers
func (v *Vertex) ScaleWithPointerReceiver(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v)
	// {3 4}

	v = Vertex{3, 4}
	v.ScaleWithValueReceiver(10)
	fmt.Println(v)
	// {3 4}
	fmt.Println(v.Abs())
	// 5

	v = Vertex{3, 4}
	v.ScaleWithPointerReceiver(10)
	fmt.Println(v)
	// {30 40}
	fmt.Println(v.Abs())
	// 50
}
