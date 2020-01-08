package main

import "fmt"

func main() {
	fmt.Println("Factorial")

	n := 5
	fmt.Println(factorial(n))

}

func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}
