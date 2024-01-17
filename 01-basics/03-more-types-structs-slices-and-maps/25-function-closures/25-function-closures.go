package main

import "fmt"

func adder() func(int) int {
	sum := 0
	// fmt.Println("outer scope", sum)
	return func(x int) int {
		sum += x
		// fmt.Println("inner scope", sum)
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

// 0 0
// 1 -2
// 3 -6
// 6 -12
// 10 -20
// 15 -30
// 21 -42
// 28 -56
// 36 -72
// 45 -90

// outer scope 0
// outer scope 0
// inner scope 0
// inner scope 0
// 0 0
// inner scope 1
// inner scope -2
// 1 -2
// inner scope 3
// inner scope -6
// 3 -6
// inner scope 6
// inner scope -12
// 6 -12
// inner scope 10
// inner scope -20
// 10 -20
// inner scope 15
// inner scope -30
// 15 -30
// inner scope 21
// inner scope -42
// 21 -42
// inner scope 28
// inner scope -56
// 28 -56
// inner scope 36
// inner scope -72
// 36 -72
// inner scope 45
// inner scope -90
// 45 -90
