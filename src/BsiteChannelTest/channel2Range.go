package main

import (
	"fmt"
	"time"
)

//Channel 通道
func main() {
	/*
	*非緩沖通道 : make(chan T)
	* 一次發送 一次接收 都是阻塞
	*緩衝通道: make(chan T,capacity)
	* 發送:緩衝區的數據滿了 才會阻塞
	* 接收:緩衝區的數據空了,才會阻塞
	 */
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
