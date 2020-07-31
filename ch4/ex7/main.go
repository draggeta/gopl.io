// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.

// Ex: modify reverse to reverse a []byte that represents a utf-8 encoded string in place.

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//!+array
	s := "Esta é uma sequência aleatória."
	a := []byte(s)
	reverseStr(a)
	fmt.Println(string(a)) // "[5 4 3 2 1 0]"
	//!-array
}

//!+rev
// reverse reverses a slice of bytes in place.
func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// reverseStr reverses first each rune and then the whole slice of bytes
func reverseStr(b []byte) {
	for i := 0; i < len(b)-1; {
		_, size := utf8.DecodeRune(b[i:])
		reverse(b[i : i+size])
		i += size
	}
	reverse(b)
}
