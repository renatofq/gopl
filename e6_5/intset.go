// This file is a derivative work of "intset"
// Original work Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Original work can be found at https://github.com/adonovan/gopl.io
// Derivative work Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 6.5 of The Go Programming Language (http://www.gopl.io/)

// e6_5 intset provides a set of integers based on a bit vector.
package e6_5

import (
	"bytes"
	"fmt"
)

const wsize = 32 << (^uint(0) >> 63)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := wordBit(x)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := wordBit(x)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll add all non-negative values of args to the set.
func (s *IntSet) AddAll(args ...int) {
	for _, x := range args {
		s.Add(x)
	}
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith sets s to the intersection of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, _ := range s.words {
		if i < len(t.words) {
			s.words[i] &= t.words[i]
		} else {
			s.words[i] &= 0
		}
	}
}

// DifferenceWith sets s to the different between s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			break
		}
	}
}

// SymmetricDifferenceWith sets s to the different between s and t.
func (s *IntSet) SymmetricDifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	s.Walk(func(n int) bool {
		if buf.Len() > len("{") {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", n)

		return false
	})

	buf.WriteByte('}')
	return buf.String()
}

// Len return the number of elements
func (s *IntSet) Len() int {
	var count int
	for _, w := range s.words {
		count += popCount(uint64(w))
	}

	return count
}

// Remove removes x element
func (s *IntSet) Remove(x int)  {
	wn, bn := wordBit(x)
	if wn < len(s.words) {
		s.words[wn] &^= 1 << bn
	}
}

// Clear remove all elements from the set
func (s *IntSet) Clear() {
	s.words = nil
}

// Copy return a copy of s
func (s *IntSet) Copy() *IntSet {
	newSet := new(IntSet)
	newSet.words = make([]uint, len(s.words))
	copy(newSet.words, s.words)
	return newSet
}

// Elems returns a slice of all elements in s
func (s *IntSet) Elems() []int {
	var elems []int

	s.Walk(func (e int) bool {
		elems = append(elems, e)
		return false
	})

	return elems
}

func (s *IntSet) Walk(f func(int)bool) {
	for i, w := range s.words {
		if w == 0 {
			continue
		}
		for j := 0; j < wsize; j++ {
			if w & (1 << uint(j)) > 0 {
				if stop := f(i*wsize + j); stop {
					return
				}
			}
		}
	}
}

func wordBit(x int) (int, uint) {
	return x/wsize, uint(x%wsize)
}

// PopCount returns the population count (number of set bits) of x.
// taken from popcount
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// can be found at https://github.com/adonovan/gopl.io
func popCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
