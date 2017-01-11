// Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 5.15 of The Go Programming Language (http://www.gopl.io/)

// e5_15 exports functions exports functions Max, Max2, Min and Min2
// that receives a varadic number of argumentts and returns the max or min
// accordingly
package e5_15

// Min return the smallest arg value, if no argument is given, Min returns 0
func Min(args ...int) int {
	var min int

	if len(args) == 0 {
		return 0
	}

	min = args[0]

	for _, n := range args[1:] {
		if n < min {
			min = n
		}
	}

	return min
}

// Max returns the largest argument, if no argument is given returns 0
func Max(args ...int) int {
	var max int

	if len(args) == 0 {
		return 0
	}

	max = args[0]

	for _, n := range args[1:] {
		if n > max {
			max = n
		}
	}

	return max
}

// Min2 returns the smallest argument value
func Min2(min int, args ...int) int {
	for _, n := range args {
		if n < min {
			min = n
		}
	}

	return min

}

// Max returns the largest argument
func Max2(max int, args ...int) int {

	for _, n := range args {
		if n > max {
			max = n
		}
	}

	return max
}
