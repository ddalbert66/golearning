package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan int, 10)

	chan2 := make(chan int, 10)

	go func() {
		for {
			select {
			case data := <-chan1:
				fmt.Println("data1 ---:", data)
			case data2 := <-chan2:
				fmt.Println("data2:", data2)
			}
		}
	}()

	go func() {
		for i := 1; i < 10; i++ {
			chan1 <- i
		}
	}()

	go func() {
		for i := 1; i < 10; i++ {
			chan2 <- i
		}
	}()

	time.Sleep(time.Second * 2)

	fmt.Println("---over")
}
