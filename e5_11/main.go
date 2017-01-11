// This file is a derivative work of "toposort"
// Original work Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Original work can be found at https://github.com/adonovan/gopl.io
// Derivative work Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 5.11 of The Go Programming Language (http://www.gopl.io/)

// e5_11 program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
	"os"
	"sort"
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
	"discrete math":         {"intro to programming", "linear algebra"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
	"intro to programming":  {},
}

func main() {
	sorted, err := topoSort(prereqs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	for i, course := range sorted {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	visiting := make(map[string]bool)
	seen := make(map[string]bool)
	var visitAll func(items []string) error

	visitAll = func(items []string) error {
		for _, item := range items {
			if visiting[item] {
				return fmt.Errorf("%s", item)
			}

			if !seen[item] {
				visiting[item] = true
				seen[item] = true
				if err := visitAll(m[item]); err != nil {
					return fmt.Errorf("%s -> %s", item, err)
				}
				visiting[item] = false
				order = append(order, item)
			}
		}
		return nil
	}

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	if err := visitAll(keys); err != nil {
		return nil, fmt.Errorf("dependency cycle in: %s", err)
	}
	return order, nil
}
