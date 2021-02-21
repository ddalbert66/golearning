package main

import "fmt"

func main() {
	fmt.Println(pipe(5, increment, increment, increment))
}

func increment(oriNum int) int {
	oriNum = oriNum + 1
	return oriNum
}

func pipe(oriNum int, funcs ...func(oriNum int) int) int {
	for _, f := range funcs {
		oriNum = f(oriNum)
	}
	return oriNum
}
