// This file is a derivative work of "tempflag"
// Original work Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Original work can be found at https://github.com/adonovan/gopl.io
// Derivative work Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 7.6 of The Go Programming Language (http://www.gopl.io/)
package main

import (
	"flag"
	"fmt"
)

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
