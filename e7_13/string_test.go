// Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 7.13 of The Go Programming Language (http://www.gopl.io/)
package e7_13

import (
	"testing"
)

// Compares two Expr to check equivalence
func equalExpr(a, b Expr) bool {
	switch aT := a.(type) {
	case Var:
		// could have used reflect.TypeOf to avoid
		// comparing the types everytime but it feels
		// kind of cheating since the book is not there
		// yet
		bT, ok := b.(Var)
		if !ok {
			return false
		}
		return aT == bT

	case literal:
		bT, ok := b.(literal)
		if !ok {
			return false
		}
		return aT == bT

	case unary:
		bT, ok := b.(unary)
		if !ok {
			return false
		}
		return aT.op == bT.op && equalExpr(aT.x, bT.x)

	case binary:
		bT, ok := b.(binary)
		if !ok {
			return false
		}
		return aT.op == bT.op && equalExpr(aT.x, bT.x) && equalExpr(aT.y, bT.y)

	case call:
		bT, ok := b.(call)
		if !ok {
			return false
		}

		if aT.fn != bT.fn || len(aT.args) != len(bT.args) {
			return false
		}

		for i := 0; i < len(aT.args); i++ {
			if !equalExpr(aT.args[i], bT.args[i]) {
				return false
			}
		}

		return true

	default:
		panic("unexpected Expr type")
	}
}

func TestPrint(t *testing.T) {
	tests := []string{
		"sqrt(A / pi)",
		"pow(x, 3) + pow(y, 3)",
		"pow(x, 3) + pow(y, 3)",
		"5 / 9 * (F - 32)",
		"5 / 9 * (F - 32)",
		"5 / 9 * (F - 32)",
		"-1 + -x",
		"-1 - x",
		"((-1)) - (((x)))",
	}

	for _, strExpr := range tests {

		expr, err := Parse(strExpr)
		if err != nil {
			t.Error(err)
			continue
		}

		printedExpr, err := Parse(expr.String())
		if err != nil {
			t.Errorf("Printed expresiion could not be parsed: %s", err)
			continue
		}

		if !equalExpr(expr, printedExpr) {
			t.Errorf("%s (printed) does not yield the same ast of %s",
				expr.String(), strExpr)
			continue
		}

		if printedExpr.String() != expr.String() {
			t.Errorf("Reprinted expr '%s' isn't equal to the printed one '%s'",
				printedExpr.String(), expr.String())
		}
	}
}
