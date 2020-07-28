// Ex: write program to print sha256 hash of input by default, 384 and 512 with flags

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	algPtr := flag.String("alg", "sha256", "enter one of sha256, sha384 or sha512")
	strPtr := flag.String("str", "", "enter a string to calculate the hash of")

	flag.Parse()

	if *strPtr == "" {
		fmt.Println("no string entered")
		return
	}

	switch *algPtr {
	case "sha256":
		fmt.Printf("%x\n", sha256.Sum256([]byte(*strPtr)))
	case "sha384":
		fmt.Printf("%x\n", sha512.Sum384([]byte(*strPtr)))
	case "sha512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(*strPtr)))
	}
}
