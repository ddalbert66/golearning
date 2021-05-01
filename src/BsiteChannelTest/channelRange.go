package main

import (
	"fmt"
	"time"
)

//Channel 通道
func main() {
	var c1 chan int
	c1 = make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			v, s := <-c1
			time.Sleep(time.Second)
			fmt.Println("打印數字", v, s)
		}

	}()

	sendData(c1)

}

func sendData(c1 chan int) {
	for i := 0; i < 10; i++ {
		c1 <- i
		fmt.Println("發送資料", i)
	}

}
