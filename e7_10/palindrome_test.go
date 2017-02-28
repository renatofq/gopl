// Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 7.10 of The Go Programming Language (http://www.gopl.io/)
package e7_10_test

import (
	"sort"
	"testing"

	"github.com/renatofq/gopl/e7_10"
)

func TestStringSlice(t *testing.T) {
	if !e7_10.IsPalindrome(sort.StringSlice([]string{"aaa"})) {
		t.Errorf("['aaa'] is palindrome")
	}

	if !e7_10.IsPalindrome(sort.StringSlice([]string{"aaa", "bbb", "aaa"})) {
		t.Errorf("['aaa', 'bbb', 'aaa'] is palindrome")
	}

	if !e7_10.IsPalindrome(sort.StringSlice([]string{"aaa", "aaa"})) {
		t.Errorf("['aaa', 'aaa'] is palindrome")
	}

	if !e7_10.IsPalindrome(
		sort.StringSlice([]string{"aaa", "bbb", "bbb", "aaa"})) {
		t.Errorf("['aaa', 'bbb' 'bbb', 'aaa'] is palindrome")
	}

	if e7_10.IsPalindrome(sort.StringSlice([]string{"aaa", "bbb"})) {
		t.Errorf("['aaa', 'bbb'] is not palindrome")
	}

	if e7_10.IsPalindrome(
		sort.StringSlice([]string{"aaa", "bbb", "ccc", "ddd"})) {
		t.Errorf("['aaa', 'bbb' 'ccc', 'ddd'] is not palindrome")
	}

	if e7_10.IsPalindrome(
		sort.StringSlice([]string{"aaa", "bbb", "ccc", "aaa"})) {
		t.Errorf("['aaa', 'bbb' 'ccc', 'aaa'] is not palindrome")
	}

	if e7_10.IsPalindrome(
		sort.StringSlice([]string{"aaa", "bbb", "bbb", "bbb"})) {
		t.Errorf("['aaa', 'bbb' 'bbb', 'bbb'] is not palindrome")
	}
}

func TestIntSlice(t *testing.T) {
	if !e7_10.IsPalindrome(sort.IntSlice([]int{1})) {
		t.Errorf("[1] is palindrome")
	}

	if !e7_10.IsPalindrome(sort.IntSlice([]int{1, 1})) {
		t.Errorf("[1, 1] is palindrome")
	}

	if !e7_10.IsPalindrome(sort.IntSlice([]int{1, 1, 1})) {
		t.Errorf("[1, 1, 1] is palindrome")
	}

	if !e7_10.IsPalindrome(sort.IntSlice([]int{1, 2, 1})) {
		t.Errorf("[1, 2, 1] is palindrome")
	}

	if !e7_10.IsPalindrome(sort.IntSlice([]int{1, 2, 2, 1})) {
		t.Errorf("[1, 2, 2, 1] is palindrome")
	}

	if e7_10.IsPalindrome(sort.IntSlice([]int{1, 2})) {
		t.Errorf("[1, 2] is not palindrome")
	}

	if e7_10.IsPalindrome(sort.IntSlice([]int{1, 2, 3, 4})) {
		t.Errorf("[1, 2, 3, 4] is not palindrome")
	}

	if e7_10.IsPalindrome(sort.IntSlice([]int{1, 2, 3, 1})) {
		t.Errorf("[1, 2, 3, 1] is not palindrome")
	}

	if e7_10.IsPalindrome(sort.IntSlice([]int{1, 2, 2, 2})) {
		t.Errorf("[1, 2, 2, 2] is not palindrome")
	}
}
