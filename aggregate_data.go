package main

import (
	"fmt"
)

type node struct {
	val int
	freq int
	left int
	right int
}
func main() {
	data := []int{1, 3, 4, 5, 2, 3, 1, 2, 2, 1,}
	freq_map := make(map[int]int)
	for i:=0; i<len(data); i++ {
		freq_map[data[i]] += 1
	}

	// for key, elem := range freq_map {

	// }
	fmt.Println(freq_map)
}

/*

Figure out how to implement a min-heap from scratch
	Make Heap --> Min-heapify --> Maintain
Would a BST be better?
Why is min-heap considered optimal?

*/