package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func main() {
	fmt.Println("Addition of two numbers =", add(42, 25))
}
