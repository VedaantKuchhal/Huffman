package main
import "fmt"

type MinHeapNode struct {
	symbol byte
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

// Calculate frequencies of different variables
func Aggregate(data string) map[byte]int {
	freq_map := make(map[byte]int)
	for i:=0; i<len(data); i++ {
		freq_map[byte(data[i])] += 1
	}
	return freq_map
}

// Swap nodes at specified indices in a given heap
func (h *MinHeap) Swap(first int, second int) {
	temp := h.heap[first]
	h.heap[first] = h.heap[second]
	h.heap[second] = temp
}

func NewMinHeap() *MinHeap {
	// Initialize an empty heap
	return &MinHeap{
		heap: []*MinHeapNode{},
		size: 0,
	}
}

func (h *MinHeap) Insert(val *MinHeapNode) {
	// Add node to end of heap and update size
	h.heap = append(h.heap, val)
	h.size++

	// Start at end of heap
	i := h.size - 1
	for i > 0 {
		// Find parent index of current node
		parent := (i - 1) / 2
		// If parent has higher frequency than child (breaking min-heap property), swap
		if h.heap[parent].freq > h.heap[i].freq {
			h.Swap(parent, i)
		}
		// Move up heap
		i = parent
	}
}

// Pop the top node
func (h *MinHeap) Pop() *MinHeapNode {
	// Return nil the heap is empty
	if h.size == 0 {
		return nil
	}
	// Save the minimum value and remove it from the heap
	min := h.heap[0]
	h.heap[0] = h.heap[h.size-1]
	h.heap = h.heap[:h.size-1]
	h.size--


	// Start at root node
	i := 0
	for i < h.size {
		// Find right and left children
		left := 2*i + 1
		right := 2*i + 2
		min := i

		// Check if left child is smaller than current
		if left < h.size && h.heap[left].freq < h.heap[min].freq {
			min = left
		}

		// Check if right child is smaller than current
		if right < h.size && h.heap[right].freq < h.heap[min].freq {
			min = right
		}
		 // Swap if min-heap property was not maintained
		if min != i {
			h.Swap(i, min)
		} else {
			break
		}
		i = min
	}

	return min
}

// Return top of heap (minimum) without deleting
func (h *MinHeap) Top() *MinHeapNode {
	return h.heap[0]
}

// Print Huffman codes
func printCodes(root *MinHeapNode, str string) {
	if root == nil {
		return
	}
	if root.symbol != '$' {
		fmt.Println(string(root.symbol), ": ", str)
	}
	printCodes(root.left, str+"0")
	printCodes(root.right, str+"1")
}

// Build Huffman Tree
func buildHuffman(agg_data map[byte]int) *MinHeap {
	// Create a new min-heap, insert all the data with corresponding frequencies
	h := NewMinHeap()
	for symbol, freq := range agg_data {
		h.Insert(&MinHeapNode{symbol, freq, nil, nil})
	}

	// Until there is one node in min-heap, build the binary tree for encoding
	for h.size != 1 {
		// Left and right children in subtree are minimum frequencies
		left := h.Pop()
		right := h.Pop()
		// New parent node is sum of frequencies of children
		tmp := &MinHeapNode{'$', left.freq + right.freq, left, right}
		// Insert parent node into min-heap
		h.Insert(tmp)
	}
	return h
}

func main(){
	data := "Helloaskdfhaslkjdhfaaaaaaaaaaaalskjhkjddd!!! !"
	agg_data := Aggregate(data)
	h := buildHuffman(agg_data)
	// Print the codes, starting from top of min-heap (only has one node left anyway)
	printCodes(h.Top(), "")
}