// Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 5.15 of The Go Programming Language (http://www.gopl.io/)

// e5_15 exports functions exports functions Max, Max2, Min and Min2
// that receives a varadic number of argumentts and returns the max or min
// accordingly
package e5_15

import (
	"testing"
)

func TestMin(t *testing.T) {
	if Min() != 0 {
		t.Errorf("Min() == 0 failed")
	}

	if Min(1) != 1 {
		t.Errorf("Min(1) == 1 failed")
	}

	if Min(1, 2, 3) != 1 {
		t.Errorf("Min(1, 2, 3) == 1 failed")
	}

	if Min(3, 2, 1) != 1 {
		t.Errorf("Min(3, 2, 1) == 1 failed")
	}

	if Min(0, -1, 1) != -1 {
		t.Errorf("Min(0, -1, 1) == 1 failed")
	}
}

func TestMax(t *testing.T) {
	if Max() != 0 {
		t.Errorf("Max(1) == 0 failed")
	}

	if Max(1) != 1 {
		t.Errorf("Max(1) == 1 failed")
	}

	if Max(1, 2, 3) != 3 {
		t.Errorf("Max(1, 2, 3) == 3 failed")
	}
}

func TestMin2(t *testing.T) {
	if Min2(1) != 1 {
		t.Errorf("Min2(1) == 1 failed")
	}

	if Min2(1, 2, 3) != 1 {
		t.Errorf("Min2(1, 2, 3) == 1 failed")
	}
}

func TestMax2(t *testing.T) {
	if Max2(1) != 1 {
		t.Errorf("Max2(1) == 1 failed")
	}

	if Max2(1, 2, 3) != 3 {
		t.Errorf("Max2(1, 2, 3) == 3 failed")
	}
}

