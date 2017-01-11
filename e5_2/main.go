// Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 5.2 of The Go Programming Language (http://www.gopl.io/)

// e5_2 counts the number of times a element appears in a html document,
// read from standard input, and prints the results
package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatalf("%s: %v\n", os.Args[0], err)
	}

	counts := make(map[string]int)
	countElements(counts, doc)
	for k, v := range counts {
		fmt.Printf("%s: %d\n", k, v)
	}
}

func countElements(counts map[string]int, n *html.Node) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode {
		counts[n.Data]++
	}

	countElements(counts, n.FirstChild)
	countElements(counts, n.NextSibling)
}
