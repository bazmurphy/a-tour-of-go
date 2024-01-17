package main

import "fmt"

func main() {
	pow := make([]int, 10)

	// If you only want the index, you can omit the second variable.
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}

	// You can skip the index or value by assigning to _
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

// 1
// 2
// 4
// 8
// 16
// 32
// 64
// 128
// 256
// 512
