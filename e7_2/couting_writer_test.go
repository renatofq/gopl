// Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 7.2 of The Go Programming Language (http://www.gopl.io/)
package e7_2_test

import (
	"testing"
	"github.com/renatofq/gopl/e7_2"
	"bytes"
)

func TestCountingWriter(t *testing.T) {
	buf := new(bytes.Buffer)
	w, nptr := e7_2.CountingWriter(buf)

	s := "1234567890"
	w.Write([]byte(s))
	if *nptr != 10 {
		t.Errorf("Number of bytes counted should be 10: %d", *nptr)
	}
	if buf.String() != "1234567890" {
		t.Errorf("Buffer shoud be equals to '%s'", s)
	}

	w.Write([]byte(""))
	if *nptr != 10 {
		t.Errorf("Number of bytes counted should be 10: %d", *nptr)
	}

	w.Write([]byte("123"))
	if *nptr != 13 {
		t.Errorf("Number of bytes counted should be 13: %d", *nptr)
	}

	w.Write([]byte(nil))
	if *nptr != 13 {
		t.Errorf("Number of bytes counted should be 13: %d", *nptr)
	}

}
