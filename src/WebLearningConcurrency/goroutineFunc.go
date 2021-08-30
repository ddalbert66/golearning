package main

import (
	"fmt"
	"time"
)

func main() {

	nums := []int{10, 11, 12, 14}
	go add(nums)

	time.Sleep(time.Second * 10)

	fmt.Println("bye ...")

}

func add(nums []int) int {
	sum := 0
	for _, num := range nums {
		fmt.Println("num:", num, ",snm:", sum)
		sum = sum + num
	}

	return sum
}
