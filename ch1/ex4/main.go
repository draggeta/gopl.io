// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Ex: modify to print the names of all files in which each duplicated line occurs

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "stdin", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts)
			f.Close()
		}
	}
	for line, filesAppeared := range counts {
		var countAppeared int
		var fileNames []string

		for fileName, n := range filesAppeared {
			countAppeared += n
			fileNames = append(fileNames, fileName)
		}

		if countAppeared > 1 {
			fmt.Printf("%d\t%s\n\t%v\n", countAppeared, line, fileNames)
		}
	}
}

func countLines(f *os.File, fileName string, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][fileName]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
