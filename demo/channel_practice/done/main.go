package main

import (
	"fmt"
	"time"
)

var count int

func doWorker(id int, ch <-chan int) {
	for n := range ch {
		count++
		fmt.Printf("Worker %d received %c\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	ch := make(chan int)
	go doWorker(id, ch)
	return ch
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)

	fmt.Println("count:", count)
}

func main() {
	chanDemo()
}
