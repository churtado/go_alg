package main

import (
	"fmt"
	"math"
	"strconv"
)

/* implementing karatsuba multiplication algorithm */

func main() {
	//a := "3141592653589793238462643383279502884197169399375105820974944592"
	//b := "2718281828459045235360287471352662497757247093699959574966967627"

	a := "5"
	b := "5"

	//fmt.Println(karatsuba(a, b))
	fmt.Println(stradd(a, b))
}

// smult converts 2 strings and multiplies them.
// It converts the result back to a string
func smult(x string, y string) string {
	_x, _ := strconv.Atoi(x)
	_y, _ := strconv.Atoi(y)
	return strconv.Itoa(_x * _y)
}

// tenexp will return a string representation of
// 10 to the power of a given number
func tenexp(n float64) string {
	return strconv.Itoa(int(math.Pow(10, n)))
}

// stradd converts 2 strings to numbers and adds them.
// It returns the result as a string
func stradd(x string, y string) string {
	// reverse both strings

	// add and carry
	return ""
}

// reverse reverses a given string and returns it
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// karatsuba implements the karatsuba algorithm for
// 2 numbers given as strings.
func karatsuba(x string, y string) string {

	// base case:
	if len(x) == 1 && len(y) == 1 {

		return smult(x, y)
	}

	// split into a, b, c, d
	n := len(x)

	a := x[0 : n/2]
	b := x[n/2 : n]
	c := y[0 : n/2]
	d := y[n/2 : n]

	// calculate powers
	f1 := tenexp(float64(n))
	f2 := tenexp(float64(n / 2))

	// recursive call
	return stradd(karatsuba(f1, karatsuba(a, c)),
		stradd(karatsuba(f2, (stradd(karatsuba(a, d), karatsuba(b, c)))),
			karatsuba(b, d)))
}
