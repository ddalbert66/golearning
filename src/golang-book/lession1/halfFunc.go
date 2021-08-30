package main

import "fmt"

func main() {
	fmt.Println(half(3))
}

func half(num int) (int, bool) {
	isEven := true
	remainder := 0
	if (num % 2) != 0 {
		isEven = false
		remainder = num % 2
	} else {
		isEven = true
		remainder = num % 2
	}
	return remainder, isEven

}
