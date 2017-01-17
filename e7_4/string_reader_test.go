// Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 7.4 of The Go Programming Language (http://www.gopl.io/)
package e7_4_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/renatofq/gopl/e7_4"
	"golang.org/x/net/html"
)

const sampleHTML = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Test page for parsing</title>
  </head>
  <body>
    <p>Some content here</p>
  </body>
</html>
`

func TestStringReaderNilP(t *testing.T) {
	var p []byte
	reader := e7_4.NewReader("ABCDE")

	n, err := reader.Read(p)
	if n != 0 || err != nil {
		t.Errorf("Error while reading to nil slice. n: %d, err: %v, p: %v",
			n, err, p)
	}

}

func TestStringReaderEmptyStr(t *testing.T) {
	var p []byte
	reader := e7_4.NewReader("")

	n, err := reader.Read(p)
	if n != 0 || err != io.EOF {
		t.Errorf("Error while reading to nil slice empty string. n: %d, err: %v, p: %v",
			n, err, p)
	}

	p = make([]byte, 5)
	n, err = reader.Read(p)
	if n != 0 || err != io.EOF {
		t.Errorf("Error while reading to nil slice empty string. n: %d, err: %v, p: %v",
			n, err, p)
	}

}

func TestStringReaderNonMultipleSize(t *testing.T) {
	reader := e7_4.NewReader("ABCDE")
	p := make([]byte, 2)

	n, err := reader.Read(p)
	if n != 2 || err != nil || bytes.Compare(p, []byte("AB")) != 0 {
		t.Errorf("Error while reading the first bytes. n: %d, err: %v, p: %v",
			n, err, p)
	}

	// Ensures p is not changed at p[n:]
	n, err = reader.Read(p)
	if n != 2 || err != nil || bytes.Compare(p, []byte("CD")) != 0 {
		t.Errorf("Error at second reading. n: %d, err: %v, p: %v",
			n, err, p)

	}

	n, err = reader.Read(p)
	if n != 1 || err != nil || bytes.Compare(p, []byte("ED")) != 0 {
		t.Errorf("Error while reading the last byte. n: %d, err: %v, p: %v",
			n, err, p)
	}

	n, err = reader.Read(p)
	if n != 0 || err != io.EOF || bytes.Compare(p, []byte("ED")) != 0 {
		t.Errorf("Error while reading after EOF. n: %d, err: %v, p: %v",
			n, err, p)

	}
}

func TestStringReaderSizeMultiple(t *testing.T) {
	reader := e7_4.NewReader("ABCD")
	p := make([]byte, 2)

	n, err := reader.Read(p)
	if n != 2 || err != nil || bytes.Compare(p, []byte("AB")) != 0 {
		t.Errorf("Error while reading the first bytes. n: %d, err: %v, p: %v",
			n, err, p)
	}

	n, err = reader.Read(p)
	if n != 2 || err != nil || bytes.Compare(p, []byte("CD")) != 0 {
		t.Errorf("Error at second reading. n: %d, err: %v, p: %v",
			n, err, p)
	}

	n, err = reader.Read(p)
	if n != 0 || err != io.EOF || bytes.Compare(p, []byte("CD")) != 0 {
		t.Errorf("Error while reading after EOF. n: %d, err: %v, p: %v",
			n, err, p)
	}
}

func TestStringReaderParseHTML(t *testing.T) {
	_, err := html.Parse(e7_4.NewReader(sampleHTML))
	if err != nil {
		t.Errorf("Fail to parse HTML from string reader: %v", err)
	}
}
