package main

import "fmt"

func main() {
	nums := []int{11, 12, 23}
	fmt.Println(sum(nums))
}

func sum(nums []int) int {
	a := 0
	for _, arg := range nums {
		a += arg
	}
	return a
}
