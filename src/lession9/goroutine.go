package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(j int) {
			for {
				a[j]++
				runtime.Gosched() //主動交出控制權
			}
		}(i)
	}

	time.Sleep(time.Millisecond)
	fmt.Println(a)

}
