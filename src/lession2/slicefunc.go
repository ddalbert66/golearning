package main

import "fmt"

func main() {
	var s = []int{1, 2, 3}
	var s2 = append(s, 1, 2)
	fmt.Println(s, s2)

	var s3 = make([]int, 2)
	copy(s3, s)
	fmt.Println(s3)
}
