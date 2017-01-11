// This file is a derivative work of "findlinks3"
// Original work Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Original work can be found at https://github.com/adonovan/gopl.io
// Derivative work Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 5.14 of The Go Programming Language (http://www.gopl.io/)

// e5_14 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"log"
	"os"
	"io/ioutil"
	"path/filepath"
)

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(path string) []string {
	fmt.Println(path)

	fileinfo, err := os.Stat(path)
	if err != nil {
		log.Print(err)
		return nil
	}

	if !fileinfo.IsDir() {
		return nil
	}

	filelist, err := ioutil.ReadDir(path)
	if err != nil {
		log.Print(err)
		return nil
	}

	var worklist []string
	for _, info := range filelist {
		worklist = append(worklist, filepath.Join(path, info.Name()))
	}

	return worklist
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
