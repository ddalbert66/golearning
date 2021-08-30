package main

import "fmt"

func main() {
	nums := []int{1, 5, 39, 22, 90, 2}
	fmt.Println(findMaxNum(nums))
}

func findMaxNum(nums []int) int {
	tempNum := 0
	for _, num := range nums {
		if tempNum < num {
			tempNum = num
		}
	}
	return tempNum
}
