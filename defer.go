package main

import "fmt"

func main() {
	defer fmt.Println("This is Second")
	defer fmt.Println("This is first to execute")
}
