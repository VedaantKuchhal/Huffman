package main
import "fmt"

type MinHeapNode struct {
	data int
	freq int

	left  *MinHeapNode
	right *MinHeapNode
}

type MinHeap struct {
	// Array to store heap elements
	heap []*MinHeapNode
	// Variable to store the size of the heap
	size int
}

func NewMinHeap() *MinHeap {
	// Initialize an empty heap
	return &MinHeap{
		heap: []*MinHeapNode{},
		size: 0,
	}
}

func (h *MinHeap) Insert (val *MinHeapNode) {
	h.heap = append(h.heap, val)
	h.size++

	i := h.size - 1
	for i > 0 {
		parent := (i - 1) / 2
		if h.heap[parent].freq > h.heap[i].freq {
			h.heap[parent], h.heap[i] = h.heap[i], h.heap[parent]
		}
		i = parent
	}
}

func (h *MinHeap) DelMin() *MinHeapNode {
	// Return if the heap is empty
	if h.size == 0 {
		return nil
	}

	// Save the minimum value and remove it from the heap
	min := h.heap[0]
	h.heap[0] = h.heap[h.size-1]
	h.heap = h.heap[:h.size-1]
	h.size--

	// Maintain the min heap property by checking the children nodes
	// and swapping the values if the current node is
	// greater than the children nodes

	i := 0
	for i < h.size {
		left := 2*i + 1
		right := 2*i + 2
		min := i

		if left < h.size && h.heap[left].freq < h.heap[min].freq {
			min = left
		}
		if right < h.size && h.heap[right].freq < h.heap[min].freq {
			min = right
		}

		if min != i {
			h.heap[i], h.heap[min] = h.heap[min], h.heap[i]
		} else {
			break
		}
		i = min
	}

	return min
}

func (h *MinHeap) Top() *MinHeapNode {
	return h.heap[0]
}

//Print Huffman codes
func printCodes(root *MinHeapNode, str string) {
	if root == nil {
		return
	}
	if root.data != '$' {
		fmt.Println(root.data, ": ", str)
	}
	printCodes(root.left, str+"0")
	printCodes(root.right, str+"1")
}

// build huffman tree
func buildHuffman(data []int, freq []int, size int) {
	var left, right *MinHeapNode
	minHeap := NewMinHeap()
	for i := 0; i < size; i++ {
		minHeap.Insert(&MinHeapNode{data[i], freq[i], nil, nil})
	}
	for minHeap.size != 1 {
		left = minHeap.Top()
		minHeap.DelMin()

		right = minHeap.Top()
		minHeap.DelMin()

		tmp := &MinHeapNode{'$', left.freq + right.freq, left, right}

		minHeap.Insert(tmp)
	}
	printCodes(minHeap.Top(), "")
}

func main(){
	arr := []int{1, 2, 3, 4, 5}
	freq := []int{10, 5, 2, 14, 15}
	size := len(arr)
	buildHuffman(arr, freq, size)
}