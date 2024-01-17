package main

import "fmt"

// A struct is a collection of fields
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}

// {1 2}

// A struct is a composite data type that groups together variables (fields) under a single name.
// It allows you to create a collection of different data types that are related to each other.
// This is similar to a class in object-oriented programming languages, but Go doesn't have classes.
// Instead, you define a struct and then create instances of that struct.
