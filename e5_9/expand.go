// Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 5.9 of The Go Programming Language (http://www.gopl.io/)

// e5_9 expands each substring "$foo" of s with f("foo"), being f a
// argument of type func(string)string
package e5_9

import (
	"bytes"
	"strings"
	"unicode"
)

// Expand replaces each substring "$foo" of s with f("foo")
func Expand(s string, f func(string) string) string {
	var buf bytes.Buffer

	i := strings.IndexRune(s, '$')
	j := 0
	for i >= 0 && j < len(s) {
		i += j
		buf.WriteString(s[j:i])

		lenOfValue := strings.IndexFunc(s[i+1:], func(r rune) bool {
			return unicode.IsSpace(r) || r == '$'
		})

		if lenOfValue < 0 {
			j = len(s)
		} else {
			j = i + lenOfValue + 1
		}

		exp := f(s[i+1 : j])
		buf.WriteString(exp)
		i = strings.IndexRune(s[j:], '$')
	}

	buf.WriteString(s[j:])

	return buf.String()
}
