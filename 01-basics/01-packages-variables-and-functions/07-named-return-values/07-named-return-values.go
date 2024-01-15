package main

import "fmt"

// the first () are the named return values
// these names should be used to document the meaning of the return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// a return statement without arguments returns the named return values
	// a "naked" return
	return
}

func main() {
	fmt.Println(split(17))
}
// 7 10