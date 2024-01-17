package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	// create an empty map of key string : value int
	// to store the words and their counts eg. "hello": 1
	wordmap := make(map[string]int)
	// fmt.Println(result)

	// split the string into words and create slice of substrings
	wordslice := strings.Fields(s)

	// loop through the slice with for key, value range slice name
	for _, value := range wordslice {
		// fmt.Println(value)

		// check if the word is already in the map
		count, ok := wordmap[value]

		// if the word is there increment
		if ok {
			wordmap[value] = count + 1
		} else {
			// if the word is not there add it
			wordmap[value] = 1
		}
	}

	// fmt.Println(result)
	// return the map
	return wordmap
}

func main() {
	// WordCount("hello how are you")
	wc.Test(WordCount)
}

// PASS
//  f("I am learning Go!") =
//   map[string]int{"Go!":1, "I":1, "am":1, "learning":1}
// PASS
//  f("The quick brown fox jumped over the lazy dog.") =
//   map[string]int{"The":1, "brown":1, "dog.":1, "fox":1, "jumped":1, "lazy":1, "over":1, "quick":1, "the":1}
// PASS
//  f("I ate a donut. Then I ate another donut.") =
//   map[string]int{"I":2, "Then":1, "a":1, "another":1, "ate":2, "donut.":2}
// PASS
//  f("A man a plan a canal panama.") =
//   map[string]int{"A":1, "a":2, "canal":1, "man":1, "panama.":1, "plan":1}
