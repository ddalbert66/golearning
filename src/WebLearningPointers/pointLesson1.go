package main

import (
	"fmt"
)

func main() {
	x := 0
	changeVal(&x)
	fmt.Println(x)

	x2 := 2
	square(&x2)
	fmt.Println(x2)

	swap(&x, &x2)
	fmt.Println(x, x2)
}

func changeVal(x *int) {
	*x = 5
}

func square(x *int) {
	*x = *x * *x
}

func swap(x3 *int, y3 *int) {
	tempx := *x3
	tempy := *y3
	*x3 = tempx
	*y3 = tempy
}
