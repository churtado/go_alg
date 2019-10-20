package main

import "fmt"

func main() {
	fmt.Println("Merge Sort on an array of positive integers")
	a := []int{4, 3, 1, 2, 5, 3}
	fmt.Println(mergeSort(a))
}

// Merge sort algorithm for an array of positive integers
func mergeSort(a []int) []int {
	// base case: array of length 1
	if len(a) == 1 {
		return a
	} else {
		b := mergeSort(a[0 : len(a)/2])
		c := mergeSort(a[len(a)/2 : len(a)])
		j := 0
		k := 0
		r := make([]int, len(a))
		for i := 0; i < len(a); {

			_b := -1
			_c := -1
			if j < len(b) {
				_b = b[j]
			}
			if k < len(c) {
				_c = c[k]
			}

			if (_b <= _c && _b > -1) || (_b != -1 && _c == -1) {
				r[i] = _b
				j++
				i++
			}

			if (_c <= _b && _c > -1) || (_c != -1 && _b == -1) {
				r[i] = _c
				k++
				i++
			}
		}
		return r
	}
}
