// This file is a derivative work of "outline2"
// Original work Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Original work can be found at https://github.com/adonovan/gopl.io
// Derivative work Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 5.8 of The Go Programming Language (http://www.gopl.io/)

// e5_8 seach and find a HTML element in a HTML document fetched from given url
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) > 2 {
		printByID(os.Args[1], os.Args[2])
	}
}

func printByID(url, id string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	elem := ElementByID(doc, id)

	fmt.Printf("%+v\n", elem)

	return nil
}

// ElementByID returns the first element whose id matches with
// id argument
func ElementByID(doc *html.Node, id string) *html.Node {

	var ret *html.Node

	forEachNode(doc, func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == "id" && attr.Val == id {
					ret = n
					return true
				}
			}
		}
		return false
	}, nil)

	return ret
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) (stop bool) {

	if pre != nil {
		stop = pre(n)
	}

	for c := n.FirstChild; c != nil && !stop; c = c.NextSibling {
		stop = forEachNode(c, pre, post)
	}

	if post != nil && !stop {
		stop = post(n)
	}

	return
}
