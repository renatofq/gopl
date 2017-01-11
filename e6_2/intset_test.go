// This file is a derivative work of "intset"
// Original work Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Original work can be found at https://github.com/adonovan/gopl.io
// Derivative work Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 6.2 of The Go Programming Language (http://www.gopl.io/)
package e6_2

import (
	"testing"
)

func TestAddHas(t *testing.T) {
	set := new(IntSet)

	set.Add(1)
	if !set.Has(1) {
		t.Errorf("Set %s should have 1", set)
	}

	set.Add(1024)
	if !set.Has(1) || !set.Has(1024) {
		t.Errorf("Set %s should have 1 and 1024", set)
	}

	set.Add(0)
	set.Add(1024)
	if !set.Has(0) || !set.Has(1) || !set.Has(1024) {
		t.Errorf("Set %s should have 0, 1 and 1024", set)
	}

	set.Add(-1)
	if set.Has(-1) || set.Len() != 3 {
		t.Errorf("Set cannot add negative numbers. %s should have 0, 1 and 1024",
			set)
	}
}

func TestRemove(t *testing.T) {
	set := new(IntSet)

	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Add(0)
	set.Add(1024)

	set.Remove(2)
	if set.Has(2) || set.Len() != 4 {
		t.Errorf("Set %s should have 0, 1, 3 and 1024", set)
	}

	set.Remove(-1)
	if set.Len() != 4 {
		t.Errorf("Set %s should have 0, 1, 3 and 1024", set)
	}

	set.Remove(1024)
	if set.Has(1024) || set.Len() != 3 {
		t.Errorf("Set %s should have 0, 1 and 3", set)
	}

	set.Remove(0)
	if set.Has(0) || set.Len() != 2 {
		t.Errorf("Set %s should have 1 and 3", set)
	}

	set.Remove(0)
	if set.Has(0) || set.Len() != 2 {
		t.Errorf("Set %s should have 1 and 3", set)
	}

	set.Remove(1)
	set.Remove(3)
	if set.Has(1) || set.Has(3) || set.Len() != 0 {
		t.Errorf("Set %s should be empty", set)
	}


}

func TestClear(t *testing.T) {
	set := new(IntSet)

	set.Add(1)
	set.Add(2)
	set.Clear()
	if set.Has(1) || set.Has(2) || set.Len() != 0 {
		t.Errorf("Set %s should be empty.", set)
	}

	set.Clear()
	if set.Has(1) || set.Has(2) || set.Len() != 0 {
		t.Errorf("Set %s should be empty.", set)
	}

}

func TestLen(t *testing.T)  {
	set := new(IntSet)

	if l := set.Len(); l != 0 {
		t.Errorf("Empty set should have o lenght. Returned: %d", l)
	}

	set.Add(1)
	set.Add(2)
	if l := set.Len(); l != 2 {
		t.Errorf("Set %s should have o Len() == 2. Returned: %d", set, l)
	}

	set.Add(2)
	if l := set.Len(); l != 2 {
		t.Errorf("Set %s should have o Len() == 2. Returned: %d", set, l)
	}

	set.Add(129)
	if l := set.Len(); l != 3 {
		t.Errorf("Set %s should have o Len() == 3. Returned: %d", set, l)
	}

	set.Remove(2)
	if l := set.Len(); l != 2 {
		t.Errorf("Set %s should have o Len() == 2. Returned: %d", set, l)
	}


	set.Remove(3)
	if l := set.Len(); l != 2 {
		t.Errorf("Set %s should have o Len() == 2. Returned: %d", set, l)
	}
}


func TestCopy(t *testing.T) {
	set := new(IntSet)

	newSet := set.Copy()
	if newSet.Len() != 0 {
		t.Errorf("NewSet %s should be empty", newSet)
	}

	set.Add(1)
	if newSet.Len() != 0 {
		t.Errorf("NewSet %s should be empty", newSet)
	}

	newSet = set.Copy()
	if !newSet.Has(1) || newSet.Len() != 1 {
		t.Errorf("NewSet %s should have 1", newSet)
	}

	newSet.Add(2)
	if set.Has(2) {
		t.Errorf("Set %s should not have 2", newSet)
	}

	set.Add(3)
	newSet = set.Copy()
	if !newSet.Has(1) || !newSet.Has(3) || newSet.Len() != 2 {
		t.Errorf("NewSet %s should have 1 and 3", newSet)
	}
}

func TestAddAll(t *testing.T) {
	set := new(IntSet)

	set.AddAll()

	set.AddAll(1)
	if !set.Has(1) {
		t.Errorf("Set %s should have 1", set)
	}

	set.AddAll(1, 1024)
	if !set.Has(1) || !set.Has(1024) {
		t.Errorf("Set %s should have 1 and 1024", set)
	}

	set.AddAll(0, 1024, -1)
	if !set.Has(0) || !set.Has(1) || !set.Has(1024) {
		t.Errorf("Set %s should have 0, 1 and 1024", set)
	}
}
