// Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 7.5 of The Go Programming Language (http://www.gopl.io/)
package e7_5_test

import (
	"io"
	"strings"
	"testing"

	"github.com/renatofq/gopl/e7_5"
)

func TestReader(t *testing.T) {
	r := e7_5.LimitReader(strings.NewReader("ABCDE"), 3)
	p := make([]byte, 2)

	n, err := r.Read(p)
	if n != 2 || err != nil {
		t.Errorf("Unexpected read size")
	}

	total := n
	n, err = r.Read(p)
	total += n
	if total != 3 || err != nil {
		t.Errorf("Off limits read. total: %d, err: %v", total, err)
	}

	n, err = r.Read(p)
	if n != 0 || err != io.EOF {
		t.Errorf("Off limits read")
	}
}
