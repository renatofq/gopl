// Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 5.9 of The Go Programming Language (http://www.gopl.io/)

// e5_9 expands each substring "$foo" of s with f("foo"), being f a
// argument of type func(string)string
package e5_9

import (
	"testing"
)

func TestExpand(t *testing.T) {

	if Expand("aaa", bar) != "aaa" {
		t.Errorf("Expected 'aaa'")
	}

	if "" != Expand("", bar) {
		t.Errorf("Expected ''")
	}

	if "bar" != Expand("$", bar) {
		t.Errorf("Expected 'bar'")
	}

	if "bar" != Expand("$foo", bar) {
		t.Errorf("Expected 'bar'")
	}

	if "aaabar" != Expand("aaa$foo", bar) {
		t.Errorf("Expected 'aaabar'")
	}

	if "aaa bar " != Expand("aaa $foo ", bar) {
		t.Errorf("Expected 'aaa bar '")
	}

	if "aaa bar" != Expand("aaa $", bar) {
		t.Errorf("Expected 'aaa bar'")
	}

	if "aaa bar " != Expand("aaa $ ", bar) {
		t.Errorf("Expected 'aaa bar '")
	}

	if "bar aaa" != Expand("$foo aaa", bar) {
		t.Errorf("Expected 'bar aaa'")
	}

	if "aaa bar aaa" != Expand("aaa $foo aaa", bar) {
		t.Errorf("Expected 'aaa bar aaa'")
	}

	if "aaa barbar aaa" != Expand("aaa $$foo aaa", bar) {
		t.Errorf("Expected 'aaa barbar aaa'")
	}

}

func bar(s string) string {
	return "bar"
}
