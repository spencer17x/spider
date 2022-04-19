package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 6, 2, 1, 9, 10, 8, 1}
	sort.Ints(nums)
	for _, v := range nums {
		fmt.Println(v)
	}
}
