// Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 7.1 of The Go Programming Language (http://www.gopl.io/)
package e7_1

import (
	"bufio"
	"bytes"
	"io"
)

type WordCounter int

type LineCounter int

// Write counts the number of words in b.
func (wc *WordCounter) Write(b []byte) (int, error) {

	n, err := splitCounter(bytes.NewBuffer(b), bufio.ScanWords)
	*wc += WordCounter(n)

	return n, err
}

// Write counts the number of lines in b.
func (lc *LineCounter) Write(b []byte) (int, error) {

	n, err := splitCounter(bytes.NewBuffer(b), bufio.ScanLines)
	*lc += LineCounter(n)

	return n, err
}

// splitCounter counts the number of tokens split can generate from r
func splitCounter(r io.Reader, split bufio.SplitFunc) (int, error) {
	var counter int
	scanner := bufio.NewScanner(r)
	scanner.Split(split)

	for scanner.Scan() {
		counter++
	}

	return counter, scanner.Err()
}
