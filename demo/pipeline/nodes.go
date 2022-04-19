package pipeline

import "sort"

func ArraySource(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func InMemorySort(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var inNums []int
		for num := range in {
			inNums = append(inNums, num)
		}
		sort.Ints(inNums)

		for _, num := range inNums {
			out <- num
		}
		close(out)
	}()

	return out
}

func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
	}()

	return out
}
