package main

import (
	"fmt"
	"time"
)

//當程式裡面有多個channel時,EX chan1,chan2 某一個時段 chan1 或者chan2 觸發了
//需要對應 chan1 or chan2 給予相對應的反饋
//使用select來監聽多個通道,當通道被觸發時(寫入數據,讀取數據,關閉通道)
//select語法與switch case 很像,但是所有分支條件都必須是通道IO


func main() {
	chan1 := make(chan int, 10)

	chan2 := make(chan int, 10)

	//啟動一個go func 監聽兩個channel
	go func() {
		for {
			select {
			case data := <-chan1:
				fmt.Println("data1 ---:", data)
			case data2 := <-chan2:
				fmt.Println("data2:", data2)
			}
		}
	}()

	go func() {
		for i := 1; i < 10; i++ {
			chan1 <- i
		}
	}()

	go func() {
		for i := 1; i < 10; i++ {
			chan2 <- i
		}
	}()

	time.Sleep(time.Second * 2)

	fmt.Println("---over")
}
