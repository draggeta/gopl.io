// Ex: write a function that counts the number of bits that are different in tw SHA256 hashes
package main

import (
	"crypto/sha256"
	"fmt"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountHash counts the number of different bits between two hashes
func PopCountHash(b, d [32]byte) int {
	var count int
	for i := 0; i < 32; i++ {
		x := b[i] ^ d[i]
		count += int(pc[x])
	}

	return count
}

func main() {
	h1 := sha256.Sum256([]byte("9"))
	h2 := sha256.Sum256([]byte("a"))

	t := PopCountHash(h1, h2)

	fmt.Println(t)
}
