package main

import (
	"fmt"
)

func fib(n int) int {
	if n == 2 {
		return 1
	} else if n == 1 {
		return 0
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	fmt.Println(fib(4))
}
