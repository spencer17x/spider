package main

import "fmt"

func channelDemo1() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Printf("get value is %d\n", v)
	}
}

func main() {
	channelDemo1()
}
