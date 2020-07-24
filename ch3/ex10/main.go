// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//

// Ex: create a non-recursive version using bytes.Buffer
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	var sm int

	n := len(s)
	if sm = n % 3; sm == 0 {
		sm = 3
	}
	buf.WriteString(s[:sm])

	for i := sm; i < n-2; i += 3 {
		fmt.Fprintf(&buf, ",%v", s[i:i+3])
	}

	return buf.String()
}

//!-
