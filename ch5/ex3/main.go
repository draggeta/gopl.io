// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc := html.NewTokenizer(os.Stdin)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "ch5/ex3: %v\n", err)
	// 	os.Exit(1)
	// }
	for {
		t := doc.Next()
		name, _ := doc.TagName()
		sName := string(name)

		if t == html.ErrorToken {
			break
		}
		if sName != "script" && sName != "style" {
			m := doc.Text()
			if len(strings.TrimSpace(string(m))) > 0 {
				fmt.Println(string(m))
			}
		}

	}

	//!-main
}
