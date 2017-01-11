// This file is a derivative work of "outline2"
// Original work Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Original work can be found at https://github.com/adonovan/gopl.io
// Derivative work Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 5.17 of The Go Programming Language (http://www.gopl.io/)

// e5_17 exports ElementsByTagName.
package e5_17

import (
	"golang.org/x/net/html"
)

// ElementsByTagName returns a slice of nodes whose tag name is equal to
// one of the name args
func ElementsByTagName(doc *html.Node, names ...string) []*html.Node {
	var res []*html.Node
	forEachNode(doc, func(n *html.Node) {
		if n.Type != html.ElementNode {
			return
		}

		for _, name := range names {
			if name == n.Data {
				res = append(res, n)
			}
		}
	}, nil)

	return res
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
