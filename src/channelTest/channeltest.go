package main

import (
	"fmt"
	"time"
)

var message chan string

func Bot() {
	msg := <-message
	fmt.Printf("Bot Print:%s\n", msg)
}

func main() {

	message = make(chan string)

	go Bot()

	message <- "first message"
	fmt.Println("first message send finish")
	message <- "second message"
	fmt.Println("second message send finish")
	time.Sleep(time.Second * 2)
	fmt.Println("end")
}
