package main

import (
	"fmt"
	"goLearning20200930/src/hello"
	"goLearning20200930/src/hello2"
	"goLearning20200930/src/stack"
)

func main() {
	var st = new(stack.Stack)
	st.Push(1)
	st.Push(2)
	st.Push(3)
	hello.Hello()
	hello2.Hello2()
	fmt.Println(st)

}
