package main

import (
	"fmt"
)

type MinHeap struct {
	A []int
	size int
	cap int
}

func NewHeap(cap int) *MinHeap {
	A := make([]int, cap)
	heap := &MinHeap{A, 0, cap}
	return heap
}

func Parent(i int) int {
	return (i - 1)/ 2
}

func LeftChild(i int) int {
	return 2*i + 1
}

func RightChild(i int) int {
	return 2*i + 2
}

func Swap(heap *MinHeap, first int, second int) {
	temp := heap.A[first]
	heap.A[first] = heap.A[second]
	heap.A[second] = temp
}

func Minheapify(heap *MinHeap, i int) {
	l := LeftChild(i)
	r := RightChild(i)

	smallest := i
	if (l <= heap.size && heap.A[l] < heap.A[i]) {
		smallest = l
	}
	if (r <= heap.size && heap.A[l] < heap.A[i]) {
		smallest = r
	}

	if smallest != i {
		Swap(heap, i, smallest)
		Minheapify(heap, smallest)
	}
}

func Insert(heap *MinHeap, v int) {
	if heap.size == heap.cap {
		panic("ERROR: reached capacity")
	}
	curr := heap.size
	heap.A[heap.size] = v
	for (curr > 0 && heap.A[Parent(curr)] > heap.A[curr]) {
		Swap(heap, curr, Parent(curr))
		curr = Parent(curr)
	}
	heap.size++
}

func ExtractMin(heap *MinHeap) int {
	if heap.size == 0 {
		return 0
	}
	last := heap.A[heap.size-1]
	temp := heap.A[0]
	heap.A[0] = last

	heap.size--
	Minheapify(heap, 0)
	return temp
}

func PrintHeap(heap *MinHeap) {
	var print func(int, string)
	print = func (i int, prefix string) {
		if i < heap.size {
			print(RightChild(i), prefix + "   ")
			fmt.Printf("%s%d : %d\n", prefix, i, heap.A[i])
			print(LeftChild(i), prefix + "   ")
		} else {
			fmt.Printf("%s-\n", prefix)
		}
	}
	fmt.Println(" /")
	print(0, " |  ")
	fmt.Println(" \\")
}

func main() {
	heap := NewHeap(10)
	Insert(heap, 8)
	Insert(heap, 13)
	Insert(heap, 8)
	Insert(heap, 3)
	Insert(heap, 23)
	fmt.Println(ExtractMin(heap))
	PrintHeap(heap)
}