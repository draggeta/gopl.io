// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ch5/ex1: %v\n", err)
		os.Exit(1)
	}
	links := []string{}
	visit(&links, doc)
	for _, link := range links {
		fmt.Println(link)
	}
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(links *[]string, n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				*links = append(*links, a.Val)
			}
		}
	}
	if n.FirstChild != nil {
		visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		visit(links, n.NextSibling)
	}
}

//!-visit

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
