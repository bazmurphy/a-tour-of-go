package main

import "fmt"

// type I interface {
// 	M()
// }

// type T struct {
// 	S string
// }

// // This method means type T implements the interface I,
// // but we don't need to explicitly declare that it does so.
// func (t T) M() {
// 	fmt.Println(t.S)
// }

// func main() {
// 	var i I = T{"hello"}
// 	i.M()
// }

// My Implicit Interface Example:

type PersonStruct struct {
	name string
}

type PersonInterface interface {
	sayHello()
	changeName(newName string)
}

// method with value receiver
// this is a copy of the struct defined above, and so will not mutate the original)
func (ps PersonStruct) sayHello() {
	fmt.Printf("Hello, %s!\n", ps.name)
}

// method with pointer receiver
// this will mutate the struct defined above
func (ps *PersonStruct) changeName(newName string) {
	ps.name = newName
}

// PersonStruct implements PersonInterface implicitly
// by having the two methods in PersonInterface

func main() {
	ps := PersonStruct{name: "Baz"}
	ps.sayHello()
	// Hello, Baz!
	ps.changeName("Bob")
	ps.sayHello()
	// Hello, Bob!
}
