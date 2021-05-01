package main

import "fmt"

//Channel 通道
func main() {
	var c1 chan bool
	c1 = make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("打印數字%d \n", i)
		}
		c1 <- true
	}()

	data := <-c1
	fmt.Print(data)
}
