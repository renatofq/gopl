// This file is a derivative work of "outline2"
// Original work Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Original work can be found at https://github.com/adonovan/gopl.io
// Derivative work Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 5.7 of The Go Programming Language (http://www.gopl.io/)

// e5_7 prints and indented version of HTML documents fetched from the
// given urls
package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var depth int

func main() {
	for _, url := range os.Args[1:] {
		prettyPrint(url)
	}
}

func prettyPrint(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func startElement(n *html.Node) {

	switch n.Type {
	case html.TextNode:
		printTextNode(n)
	case html.DocumentNode:
		depth--
	case html.ElementNode:
		printElementNode(n)
	case html.CommentNode:
		printCommentNode(n)
	case html.DoctypeNode:
		printDoctypeNode(n)
	}

	depth++
}

func printTextNode(n *html.Node) {
	if data := strings.TrimSpace(n.Data); len(data) > 0 {
		fmt.Printf("%*s%s\n", depth*2, "", data)
	}
}

// printElementNode prints a HTML node
func printElementNode(n *html.Node) {
	var attrBuf bytes.Buffer

	for _, attr := range n.Attr {
		fmt.Fprintf(&attrBuf, ` %s="%s"`, attr.Key, attr.Val)
	}

	if n.FirstChild == nil {
		fmt.Printf("%*s<%s%s/>\n", depth*2, "", n.Data, attrBuf.String())
	} else {
		fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, attrBuf.String())
	}
}

func printCommentNode(n *html.Node) {
	if data := strings.TrimSpace(n.Data); len(data) > 0 {
		fmt.Printf("%*s<!-- %s -->\n", depth*2, "", data)
	}
}

func printDoctypeNode(n *html.Node) {
	fmt.Printf("<!DOCTYPE %s>\n", n.Data)
}

func endElement(n *html.Node) {
	depth--

	if n.Type == html.ElementNode && n.FirstChild != nil {
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
