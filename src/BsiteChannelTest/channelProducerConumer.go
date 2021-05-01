package main

import (
	"fmt"
	"time"
)

func main() {

	chantest := make(chan int, 5)
	go produce(chantest)
	go consumer(chantest)

	time.Sleep(time.Second * 10)

}

func produce(c1 chan int) {
	for i := 1; i < 10; i++ {
		c1 <- i
		fmt.Println("生產:", i)
	}
}

func consumer(c1 <-chan int) {
	for v := range c1 {

		fmt.Println("消費:", v)
	}
}
