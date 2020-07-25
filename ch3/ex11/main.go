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

// Ex: handle ints and floats and optional signs

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", commaDecimals(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaDecimals(s string) string {
	var dec string

	if dot := strings.LastIndex(s, "."); dot != -1 {
		dec = s[dot+1:]
		s = s[:dot]
	}

	s = comma(s)

	if dec != "" {
		return fmt.Sprintf("%s.%s", s, dec)
	}
	return s
}

//!-
