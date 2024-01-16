package main

import "fmt"

func main() {
	sum := 0
	// there are no surrounding parentheses ()
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
