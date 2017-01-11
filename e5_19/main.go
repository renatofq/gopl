// Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 5.19 of The Go Programming Language (http://www.gopl.io/)
package e5_19

// WithoutReturn returns 1 (non-zero value) without a return statement
func WithoutReturn() (ret int) {
	defer func() {
		if p := recover(); p != nil {
			ret = 1
		}
	}()

	panic(1)
}
