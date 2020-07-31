// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.

// Ex: rewrite reverse to use an array pointer instead of a slice

package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "strconv"
	// "strings"
)

func main() {
	//!+array
	a := [...]int{0, 1, 2, 3, 4, 5}
	// Doing this seems to really be bad practice.
	// There is also seems to be no way to pass a variable length array
	reverse(&a)
	fmt.Println(a) // "[5 4 3 2 1 0]"
	//!-array
}

//!+rev
// reverse reverses a slice of ints in place.
func reverse(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//!-rev
