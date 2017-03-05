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

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		want string
	}{
		{"pow(12, 3) + pow(1, 3)", "1729"},
		{"pow(9, 3) + pow(10, 3)", "1729"},
		{"5 / 9 * ((-40) - 32)", "-40"},
		{"5 / 9 * (32 - 32)", "0"},
		{"5 / 9 * (212 - 32)", "100"},
	}
	var prevExpr string
	for _, test := range tests {
		// Print expr only when it changes.
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval())
		fmt.Printf("=> %s\n", got)
		if got != test.want {
			t.Errorf("%s.Eval() = %q, want %q\n",
				test.expr, got, test.want)
		}
	}
}

func TestErrors(t *testing.T) {
	for _, test := range []struct{ expr, wantErr string }{
		{"10 % 2", "unexpected '%'"},
		{"math.Pi", "unexpected '.'"},
		{"!true", "unexpected '!'"},
		{`"hello"`, "unexpected '\"'"},
		{"log(10)", `unknown function "log"`},
		{"sqrt(1, 2)", "call to sqrt has 2 args, want 1"},
	} {
		expr, err := Parse(test.expr)
		if err == nil {
			err = expr.Check()
			if err == nil {
				t.Errorf("unexpected success: %s", test.expr)
				continue
			}
		}
		fmt.Printf("%-20s%v\n", test.expr, err) // (for book)
		if err.Error() != test.wantErr {
			t.Errorf("got error %s, want %s", err, test.wantErr)
		}
	}
}
