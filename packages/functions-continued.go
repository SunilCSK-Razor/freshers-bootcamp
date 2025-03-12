package main

import "fmt"

func addV2(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("Addition of two numbers =", addV2(42, 28))
}
