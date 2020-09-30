package main

import (
	"fmt"
	"stack"
)

func main() {
	var st = new(stack.Stack)
	st.Push(1)
	st.Push(2)
	st.Push(3)

	fmt.Println(st)
}
