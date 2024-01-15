package main

// individual import statements
// import "fmt"
// import "math"

// factored import statement
import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
// Now you have 2.6457513110645907 problems.