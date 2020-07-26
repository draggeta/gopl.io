// Ex: write a function that reports if two strings are anagrams

package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please specify two strings.")
		return
	}

	w1, w2 := os.Args[1], os.Args[2]
	if len([]rune(w1)) != len([]rune(w2)) {
		fmt.Println("Not anagrams")
		return
	}

	w1Count := make(map[rune]int)
	w2Count := make(map[rune]int)

	for _, r := range w1 {
		w1Count[r]++
	}

	for _, r := range w2 {
		w2Count[r]++
	}

	// could compare each key, but this was more neat (but slower)
	if reflect.DeepEqual(w1Count, w2Count) {
		fmt.Println("Are anagrams")
	} else {
		fmt.Println("Not anagrams")
	}

}
