// Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 7.2 of The Go Programming Language (http://www.gopl.io/)
package e7_2

import (
	"io"
)

type writerCounter struct {
	w io.Writer
	c int64
}

// CountingWriter returns a writer that writes to w and counts the number
// of bytes written
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	wc := writerCounter{w, 0}
	return &wc, &(wc.c)
}

func (wc *writerCounter) Write(b []byte) (int, error) {
	n, err := wc.w.Write(b)
	wc.c += int64(n)
	return n, err
}
