package main

import "fmt"

func main() {
	fmt.Println("Fibonacci")

	fmt.Println(fibonacci(4))
	fmt.Fprintln("!!!")
}

func fibonacci(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	return n + fibonacci(n-1) + fibonacci(n-2)
}

//asdf
