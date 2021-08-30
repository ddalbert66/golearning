package main

import "fmt"

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)

	}()
	defer fmt.Println("2")
	panic("PANIC")
}
