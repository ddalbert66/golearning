package main

import "fmt"

func main() {
	nums := [10]int{1, 2, 3, 4, 5}
	chanCheck := make(chan int, 10)
	isOver := make(chan int)
	go func() {

		for i := 0; i < len(nums); i++ {
			chanCheck <- i
			fmt.Println(nums[i], "start---")

		}
	}()

	go func() {
		for _, v := range nums {
			result := <-chanCheck

			if result == 9 {
				isOver <- 1
			}
			fmt.Println(v, "----reciver")
		}
	}()

	<-isOver
	fmt.Println("OVER")
}
