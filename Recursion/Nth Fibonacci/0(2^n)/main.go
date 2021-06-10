//Nth fibonacci number formula
//F(n) = F(n-1)+F(n-2)

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
