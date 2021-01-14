package main

import (
	"fmt"
	"sync"
)

func doWork(id int, c chan int,wg *sync.WaitGroup) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
		go func() {
			wg.Done()
		}()

	}
}

type worker struct {
	in   chan int
	wg *sync.WaitGroup
}

func createWorker(id int,wg *sync.WaitGroup) worker {
	w := worker{
		in:   make(chan int),
		wg:wg,
	}
	go doWork(id, w.in, wg)
	return w
}

func chanDemo() {
	var workers [10]worker
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	


	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()
}

func main() {
	chanDemo()

}
