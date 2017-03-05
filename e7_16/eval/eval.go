// This file is a adapted of "eval" to meet exercises needs
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Original work can be found at https://github.com/adonovan/gopl.io
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 7.16 of The Go Programming Language (http://www.gopl.io/)
package eval

import (
	"fmt"
	"math"
)


func (l literal) Eval() float64 {
	return float64(l)
}

func (u unary) Eval() float64 {
	switch u.op {
	case '+':
		return +u.x.Eval()
	case '-':
		return -u.x.Eval()
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval() float64 {
	switch b.op {
	case '+':
		return b.x.Eval() + b.y.Eval()
	case '-':
		return b.x.Eval() - b.y.Eval()
	case '*':
		return b.x.Eval() * b.y.Eval()
	case '/':
		return b.x.Eval() / b.y.Eval()
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval() float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(), c.args[1].Eval())
	case "sin":
		return math.Sin(c.args[0].Eval())
	case "sqrt":
		return math.Sqrt(c.args[0].Eval())
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}
