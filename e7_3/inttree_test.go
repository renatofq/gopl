// This file is a derivative work of "treesort"
// Original work Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Original work can be found at https://github.com/adonovan/gopl.io
// Derivative work Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 7.3 of The Go Programming Language (http://www.gopl.io/)
package e7_3_test

import (
	"testing"

	"github.com/renatofq/gopl/e7_3"
)

func TestSort(t *testing.T) {

	tree := e7_3.Add(nil, 2)
	tree = e7_3.Add(tree, 3)
	tree = e7_3.Add(tree, 1)

	s := tree.String()
	expected := "[1, 2, 3, ]"
	if s == expected {
		t.Errorf("String() returned %s. expected: %s", s, expected)
	}
}
