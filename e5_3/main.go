// Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 5.3 of The Go Programming Language (http://www.gopl.io/)

// e5_3 prints the text elements of a html document read from standard input
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

	printText(doc, nil)
}

func printText(n *html.Node, p *html.Node) {
	if n == nil {
		return
	}

	if n.Type == html.TextNode &&
		(p == nil || (p.Data != "script" && p.Data != "style")) {
		fmt.Println(n.Data)
	}

	printText(n.FirstChild, n)
	printText(n.NextSibling, p)
}
