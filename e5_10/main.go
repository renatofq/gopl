// This file is a derivative work of "toposort"
// Original work Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Original work can be found at https://github.com/adonovan/gopl.io
// Derivative work Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 5.10 of The Go Programming Language (http://www.gopl.io/)

// e5_10 program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}


func main() {
	for course, deps := range topoSort(prereqs) {
		fmt.Printf("%s: %v\n", course, deps)
	}
}

func topoSort(m map[string][]string) map[string][]string {

	var visit func([]string)[]string
	result := make(map[string][]string)
	seen := make(map[string]bool)

	visit = func(items []string, ) []string {
		var deps []string
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				deps = append(deps, visit(m[item])...)
				deps = append(deps, item)
			}
		}
		return deps
	}

	for k, v := range m {
		result[k] = visit(v)
		seen = make(map[string]bool)
	}

	return result
}
