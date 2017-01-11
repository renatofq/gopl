// Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 5.16 of The Go Programming Language (http://www.gopl.io/)
package e5_16_test

import (
	"testing"
	"github.com/renatofq/gopl/e5_16"
)

func TestEmptyJoin(t *testing.T) {
	res := e5_16.Join(",")

	if res != "" {
		t.Errorf("Empty joing should return a empty string. returned: %s", res)
	}
}

func TestSingleJoin(t *testing.T) {
	res := e5_16.Join(",", "foo")

	if res != "foo" {
		t.Errorf("Empty joing should return a empty string. returned: %s", res)
	}
}

func TestJoin(t *testing.T) {
	res := e5_16.Join(",", "foo", "bar", "baz")

	if res != "foo,bar,baz" {
		t.Errorf("Empty joing should return a empty string. returned: %s", res)
	}
}
