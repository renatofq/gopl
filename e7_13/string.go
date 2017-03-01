// Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 7.13 of The Go Programming Language (http://www.gopl.io/)
package e7_13

import (
	"bytes"
	"fmt"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%g", float64(l))
}

func (u unary) String() string {
	return fmt.Sprintf("(%c%s)", u.op, u.x.String())
}

func (b binary) String() string {
	return fmt.Sprintf("(%s %c %s)", b.x.String(), b.op, b.y.String())
}

func (c call) String() string {
	var buf bytes.Buffer
	for i, arg := range c.args {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(arg.String())
	}
	return fmt.Sprintf("%s(%s)", c.fn, buf.String())
}
