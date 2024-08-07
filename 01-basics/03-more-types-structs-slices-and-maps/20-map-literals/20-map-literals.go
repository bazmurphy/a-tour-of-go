package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// Map literals are like struct literals, but the keys are required.
var m = map[string]Vertex{
	// "Bell Labs" is the key
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	// "Google" is the key
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}

// map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
