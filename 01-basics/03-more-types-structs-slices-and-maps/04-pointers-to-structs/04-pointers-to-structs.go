package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	// make p a pointer to struct v with the & operator
	p := &v
	// change the X field on the struct v using the p.field pointer syntax
	p.X = 1e9
	fmt.Println(v)
}

// {1000000000 2}
