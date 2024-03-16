package maxheap

type MaxHeap struct {
	arr []int
}

func (h *MaxHeap) Insert(val int) {
	h.arr = append(h.arr, val)
	h.heapifyUp(len(h.arr) - 1)
}

func (h *MaxHeap) Extract() int {
	extracted := h.arr[0]
	if len(h.arr) == 0 {
		return -1
	}
	lastIndex := len(h.arr) - 1
	h.arr[0] = h.arr[lastIndex]
	h.arr = h.arr[:lastIndex]
	h.heapifyDown(0)
	return extracted
}

func (h *MaxHeap) heapifyUp(index int) {
	for h.arr[parent(index)] < h.arr[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) heapifyDown(index int) {
	last := len(h.arr) - 1
	leftIndex, rightIndex := left(index), right(index)
	childToCompare := 0

	for leftIndex <= last {
		if leftIndex == last {
			childToCompare = leftIndex
		} else if h.arr[leftIndex] > h.arr[rightIndex] {
			childToCompare = leftIndex
		} else {
			childToCompare = rightIndex
		}
		if h.arr[index] < h.arr[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			leftIndex = left(index)
			rightIndex = right(index)
		} else {
			return
		}
	}
}

func (h *MaxHeap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}
