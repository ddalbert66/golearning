package main

import "fmt"

func main() {
	// panic("PPPPPPpanic")
	// str := recover()
	// fmt.Println(str)

	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	panic("pannic")
}
