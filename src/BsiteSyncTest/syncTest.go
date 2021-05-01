package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mutex sync.Mutex
var wg sync.WaitGroup

//全局變量 表示票
var ticket = 10

func main() {
	wg.Add(4)
	go saleTickets("售票口1")
	go saleTickets("售票口2")
	go saleTickets("售票口3")
	go saleTickets("售票口4")

	wg.Wait()
}

func saleTickets(name string) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	for {
		mutex.Lock()
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			fmt.Println(name, "售出:", ticket)
			ticket--
		} else {
			mutex.Unlock()
			fmt.Println(name, "售鑿,沒有票了..")
			break
		}
		mutex.Unlock()
	}

}
