// This file is a adapted of "eval" to meet exercises needs
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Original work can be found at https://github.com/adonovan/gopl.io
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 7.16 of The Go Programming Language (http://www.gopl.io/)
package eval

import (
	"fmt"
	"strings"
)

func (literal) Check() error {
	return nil
}

func (u unary) Check() error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary op %q", u.op)
	}
	return u.x.Check()
}

func (b binary) Check() error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unexpected binary op %q", b.op)
	}
	if err := b.x.Check(); err != nil {
		return err
	}
	return b.y.Check()
}

func (c call) Check() error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}
	if len(c.args) != arity {
		return fmt.Errorf("call to %s has %d args, want %d",
			c.fn, len(c.args), arity)
	}
	for _, arg := range c.args {
		if err := arg.Check(); err != nil {
			return err
		}
	}
	return nil
}

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}
