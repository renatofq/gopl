// This file is a derivative work of "treesort"
// Original work Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Original work can be found at https://github.com/adonovan/gopl.io
// Derivative work Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 7.3 of The Go Programming Language (http://www.gopl.io/)
package e7_3

import (
	"bytes"
	"fmt"
)

type Tree struct {
	value       int
	left, right *Tree
}

func Add(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{value: v}
	}

	if v < t.value {
		t.left = Add(t.left, v)
	} else {
		t.right = Add(t.right, v)
	}

	return t
}

func Walk(t *Tree, f func(int)) {
	if t == nil {
		return
	}

	f(t.value)
	Walk(t.left, f)
	Walk(t.right, f)
}

// String returns a string with the sequence of values in the tree
func (t *Tree) String() string {
	buf := new(bytes.Buffer)

	buf.WriteString("[")
	Walk(t, func(v int) {
		buf.WriteString(fmt.Sprintf("%d, ", v))
	})
	buf.WriteString("]")

	return buf.String()
}
