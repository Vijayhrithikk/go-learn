package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
func main() {
	defer fmt.Println("This is Second")
	defer fmt.Println("This is first to execute")
	fmt.Println(fib(5))
}
