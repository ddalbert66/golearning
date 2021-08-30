package main

import "fmt"

func main() {
	fmt.Println(rabbitQuestion(9))
}

func rabbitQuestion(n uint) uint {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	n = n - 1
	return rabbitQuestion(n) + rabbitQuestion(n-1)

}
