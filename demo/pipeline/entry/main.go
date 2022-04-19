package main

import (
	"fmt"
	"spider/demo/pipeline"
)

func main() {
	nums1 := []int{3, 2, 6, 7, 4}
	nums2 := []int{7, 4, 0, 3, 2, 13, 8}
	ch := pipeline.Merge(
		pipeline.InMemorySort(
			pipeline.ArraySource(nums1...),
		),
		pipeline.InMemorySort(
			pipeline.ArraySource(nums2...),
		),
	)
	for n := range ch {
		fmt.Printf("get value is %v\n", n)
	}
}
