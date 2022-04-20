package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

var count int

func doWorker(id int, ch <-chan int, wg *sync.WaitGroup) {
	for n := range ch {
		count++
		fmt.Printf("Worker %d received %c\n", id, n)
		wg.Done()
	}
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWorker(id, w.in, wg)
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

	//time.Sleep(time.Millisecond)

	fmt.Println("count:", count)
}

func main() {
	chanDemo()
}
