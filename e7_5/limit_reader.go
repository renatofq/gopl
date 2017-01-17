// Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 7.5 of The Go Programming Language (http://www.gopl.io/)
package e7_5

import (
	"io"
)

type limitReader struct {
	reader io.Reader
	limit  int64
	pos    int64
}

// LimitReader reads from r up to the limit of n bytes, when it starts
// roporting end-of-file
func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{reader: r, limit: n}
}

func (lr *limitReader) Read(p []byte) (n int, err error) {
	var readLen int64

	// avoids to change the the internal reader when it comes to
	// handling nil == p and whether it returns err == io.EOF and
	// n > 0 at the same call
	if lr.pos == lr.limit && len(p) > 0 {
		return 0, io.EOF
	}

	if lr.limit < lr.pos + int64(len(p)) {
		readLen = lr.limit - lr.pos
	} else {
		readLen = int64(len(p))
	}

	n, err = lr.reader.Read(p[:readLen])
	lr.pos += int64(n)

	return n, err
}
