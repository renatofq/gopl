// This file is a adapted of "eval" to meet exercises needs
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Original work can be found at https://github.com/adonovan/gopl.io
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 7.16 of The Go Programming Language (http://www.gopl.io/)
package eval

import (
	"fmt"
	"testing"
)

func TestCoverage(t *testing.T) {
	var tests = []struct {
		input string
		want  string // expected error from Parse/Check or result from Eval
	}{
		{"10 % 2", "unexpected '%'"},
		{"!true", "unexpected '!'"},
		{"log(10)", `unknown function "log"`},
		{"sqrt(1, 2)", "call to sqrt has 2 args, want 1"},
		{"pow(9, 3) + pow(10, 3)", "1729"},
		{"5 / 9 * ((-40) - 32)", "-40"},
	}

	for _, test := range tests {
		expr, err := Parse(test.input)
		if err == nil {
			err = expr.Check()
		}
		if err != nil {
			if err.Error() != test.want {
				t.Errorf("%s: got %q, want %q", test.input, err, test.want)
			}
			continue
		}

		got := fmt.Sprintf("%.6g", expr.Eval())
		if got != test.want {
			t.Errorf("%s => %s, want %s",
				test.input, got, test.want)
		}
	}
}
