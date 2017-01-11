// Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 4.9 of The Go Programming Language (http://www.gopl.io/)

// e4_9 prints the number of times each word appeared.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		counts[strings.ToLower(scanner.Text())]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "reading error: %v\n", err)
	}

	for k, v := range counts {
		fmt.Printf("%s: %d\n", k, v)
	}
}
