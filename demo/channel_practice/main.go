package main

import (
	"fmt"
	"time"
)

func channelDemo() {
	c := make(chan int)
	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()
	c <- 1
	c <- 2
	time.Sleep(time.Microsecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go func() {
		for {
			fmt.Printf("get value is %d\n", <-c)
		}
	}()
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	time.Sleep(time.Millisecond)
}

func worker(id int, c <-chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func channelClose() {
	channels := make([]chan<- int, 10)
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
}

func main() {
	//channelDemo()
	//bufferedChannel()
	channelClose()
}
