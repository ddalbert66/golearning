package main

import (
	"fmt"
	"time"
)

func chanIn(c chan string) {
	for i := 0; ; i++ {
		c <- "package"
	}

}

func chanOut(c chan string) {
	for {
		str := <-c
		fmt.Println(str)
		time.Sleep(time.Second * 1)
	}

}

func main() {

	var c chan string = make(chan string)
	go chanIn(c)
	go chanOut(c)
	var s string
	fmt.Scanln(&s)

}
