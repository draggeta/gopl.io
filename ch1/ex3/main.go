// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Ex: measure the running time between the two different ways.

package main

import (
	"fmt"
	"strings"
	"time"
)

var s = []string{"some", "random", "text", "here", "this", "should", "be", "enough", "!!!!"}
var sep = " "

func main() {
	start := time.Now()
	fmt.Println(strings.Join(s, sep))
	finish := time.Since(start).Seconds()
	fmt.Printf("Join: %v\n", finish)

	start = time.Now()
	lstr := s[0]
	for i := 1; i < len(s); i++ {
		lstr += sep + s[i]
	}
	fmt.Println(lstr)
	finish = time.Since(start).Seconds()
	fmt.Printf("Loop: %v\n", finish)

}
