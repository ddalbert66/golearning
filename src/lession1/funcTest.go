package main

import "fmt"

func eval(a, b int, op string) int {
	var result int
	switch {
	case op == "/":
		result = a / b
	case op == "*":
		result = a * b
	case op == "+":
		result = a + b
	case op == "-":
		result = a - b
	default:
		panic("has Error" + op)
	}
	return result
}

func main() {
	fmt.Println(eval(2, 3, "+"))
}
