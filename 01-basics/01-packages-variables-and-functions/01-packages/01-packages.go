// Every Go program is made up of packages.

// Programs start running in package main
package main

// This program is using the packages with import paths "fmt" and "math/rand"
import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
// My favorite number is 9