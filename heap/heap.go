package heap

import "fmt"

type Heap struct {
	// Checks if heap property is valid for given keys
	comparator func(int, int) bool
	// Values stored in heap
	values []int
}

// Creates new heap backed by int slice of given size
// as capacity. It uses given comp function to check
// heap property is satisfied
func NewHeap(comp func(int, int) bool, size int) *Heap {
	return &Heap{
		comparator: comp,
		values:     make([]int, 0, size),
	}
}

// Peek provides first value in heap order
// without removing value from heap
func (h *Heap) Peek() (int, error) {
	size := len(h.values)
	if size == 0 {
		return 0, fmt.Errorf("empty heap : %d elements", size)
	}

	return h.values[0], nil
}

// Set adds given value to heap
func (h *Heap) Set(val int) {
	h.values = append(h.values, val)

	h.heapifyUp()
}

// Get removes first value in heap order
// from heap and returns it
func (h *Heap) Get() int {
	size := len(h.values) - 1

	val := h.values[0]
	h.values[0] = h.values[size]
	h.values = h.values[:size]

	h.heapifyDown()

	return val
}

func (h *Heap) heapifyUp() {
	size := len(h.values)
	if size < 2 {
		return
	}

	child := size - 1
	parent := child / 2

	for parent != child {
		if h.comparator(h.values[parent], h.values[child]) {
			return
		}
		h.values[child], h.values[parent] = h.values[parent], h.values[child]
		child = parent
		parent = child / 2
	}
}

func (h *Heap) heapifyDown() {
	size := len(h.values)
	if size < 2 {
		return
	}

	parent := 0
	lChild := (parent * 2) + 1
	rChild := lChild + 1

	for lChild < size {
		child := lChild
		if rChild < size {
			if h.comparator(h.values[rChild], h.values[lChild]) {
				child = rChild
			}
		}

		if h.comparator(h.values[parent], h.values[child]) {
			return
		}
		h.values[child], h.values[parent] = h.values[parent], h.values[child]
		parent = child
		lChild = (parent * 2) + 1
		rChild = lChild + 1

	}
}
