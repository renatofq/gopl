// Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 7.4 of The Go Programming Language (http://www.gopl.io/)
package e7_4

import (
	"io"
)

type stringReader struct {
	pos int
	str string
}

func NewReader(str string) io.Reader {
	return &stringReader{str: str, pos: 0}
}

func (sr *stringReader) Read(p []byte) (n int, err error) {
	n = copy(p, sr.str[sr.pos:])
	sr.pos += n

	// mimics the behavior of strings.NewReader reader
	if n == 0 && sr.pos == len(sr.str) {
		err = io.EOF
	}

	// returnig n > 0 and err == io.EOF at the same call
	// if sr.pos == len(sr.str) && n < len(p) {
	// 	err = io.EOF
	// }

	return
}
