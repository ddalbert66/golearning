package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c1 <- "c1Msg"
			time.Sleep(time.Second)
		}

	}()

	go func() {
		for i := 0; ; i++ {
			c2 <- "c2Msg"
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			select {
			case str1 := <-c1:
				fmt.Println(str1)
			case str2 := <-c2:
				fmt.Println(str2)
			}
		}

	}()

	var input string
	fmt.Scanln(&input)

}
