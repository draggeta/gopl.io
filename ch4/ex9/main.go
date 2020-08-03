package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := make(map[string]int)
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Please enter one file path\n")
		return
	}
	file := os.Args[1]
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
		return
	}
	countWords(f, count)

	for k, v := range count {
		fmt.Printf("%d\t%s\n", v, k)
	}

}

func countWords(f *os.File, c map[string]int) {
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		c[input.Text()]++
	}
}
