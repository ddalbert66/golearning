package main

import (
	"fmt"
	"math"
)

func Trangle() {
	var a, b int = 3, 4
	var c int = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func Enum() {
	const (
		java = iota
		script
		golangxx
	)

	fmt.Println(java, golangxx, script)
}

func main() {
	Trangle()
	Enum()
}
