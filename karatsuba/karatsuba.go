package main

import (
	"fmt"
	"strconv"
)

/* implementing karatsuba multiplication algorithm */

func main() {

	// we define 2 integers

	// a := "3141592653589793238462643383279502884197169399375105820974944592"
	// b := "2718281828459045235360287471352662497757247093699959574966967627"

	// https://www.dcode.fr/big-numbers-multiplication
	// Correct:       8539734222673567065463550869546574495034888535765114961879601127067743044893204848617875072216249073013374895871952806582723184
	// Mine:          8539734222673567065463550869546574495034888535765114961879601127067743044893204848617875072216249073013374895871952806582723184

	// fmt.Println(stradd("99", "99"))
	// fmt.Println(stradd("39618900", "1404"))

	a := "1"
	b := "1"
	// fmt.Println(a, "*", b, "=", karatsuba(a, b))

	// a = "5"
	// b = "5"
	// fmt.Println(a, "*", b, "=", karatsuba(a, b))

	// a = "20"
	// b = "50"
	// fmt.Println(a, "*", b, "=", karatsuba(a, b)) // 1000

	// a = "35"
	// b = "50"
	// fmt.Println(a, "*", b, "=", karatsuba(a, b)) // 1750

	// a = "3000"
	// b = "9000"
	// fmt.Println(a, "*", b, "=", karatsuba(a, b)) // 27000000

	// a = "4527"
	// b = "8752"
	// fmt.Println(a, "*", b, "=", karatsuba(a, b)) // 39620304

	a = "3141592653589793238462643383279502884197169399375105820974944592"
	b = "2718281828459045235360287471352662497757247093699959574966967627"
	fmt.Println(a, "*", b, "=", karatsuba(a, b)) // 39620304

}

// karatsuba implements the karatsuba algorithm for
// 2 numbers given as strings.
func karatsuba(x string, y string) string {

	// base case: 2 single-digit integers
	if len(x) == 1 && len(y) == 1 {
		_x, _ := strconv.Atoi(x)
		_y, _ := strconv.Atoi(y)
		return strconv.Itoa(_x * _y)
	}

	// recursive case: split into 4 substrings
	// 4 substrings will be a, b, c, d
	// for simplicity's sake, we're assuming len(x) = len(y) and they're even
	n := len(x)

	a := x[0 : n/2]
	b := x[n/2 : n]
	c := y[0 : n/2]
	d := y[n/2 : n]

	// recursive calls
	ac := karatsuba(a, c)
	ad := karatsuba(a, d)
	bc := karatsuba(b, c)
	bd := karatsuba(b, d)

	s1 := tenmult(n, ac)

	s2 := stradd(ad, bc)
	s3 := tenmult(n/2, s2)

	s4 := stradd(s1, s3)

	s5 := stradd(s4, bd)

	return s5 //bd
}

// tenmult will concatenate n zeroes to x
func tenmult(n int, x string) string {
	for i := 0; i < n; i++ {
		x = x + "0"
	}
	return x
}

// stradd converts 2 strings to numbers and adds them.
// It returns the result as a string
func stradd(x string, y string) string {

	// pad the smaller number with zeros to the left
	if len(x) > len(y) {
		for i := len(y); i < len(x); i++ {
			y = "0" + y
		}
	} else {
		for i := len(x); i < len(y); i++ {
			x = "0" + x
		}
	}

	_x := reverse(x)
	_y := reverse(y)

	z := ""
	c := 0
	pc := 0
	s := 0
	i := 0

	// add the shorter string into the longer
	for i = 0; i < len(_y); i = i + 1 {
		a := _x[i]
		b := _y[i]
		s, c = stradd1(a, b)
		if (s + pc) < 10 {
			z = z + strconv.Itoa(s+pc)
		} else {
			t := s + pc
			z = z + strconv.Itoa(t%10)
			c = t / 10
		}

		pc = c
	}

	if c != 0 {
		z = z + strconv.Itoa(c)
	}

	z = reverse(z)
	return z
}

// reverse returns the reverse of a given string
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// stradd1 returns the addition of 2 digits and the carry
func stradd1(x byte, y byte) (int, int) {
	// convert digits to integers
	_x, _ := strconv.Atoi(string(x))
	_y, _ := strconv.Atoi(string(y))
	sum := _x + _y

	if sum < 10 {
		return sum, 0
	}

	return sum % 10, sum / 10
}

// smult converts 2 strings and multiplies them.
// It converts the result back to a string
func smult(x string, y string) string {
	_x, _ := strconv.Atoi(x)
	_y, _ := strconv.Atoi(y)
	return strconv.Itoa(_x * _y)
}
