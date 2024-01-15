package main

import "fmt"

// define a list of variables at the package level
// the type is last
var c, python, java bool

func main() {
	// define a variable at the function level
	var i int
	fmt.Println(i, c, python, java)
}
// 0 false false false