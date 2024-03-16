package examples

import (
	minheap "data-structures/types/minHeap"
	"fmt"
)

func MinHeapExample() {
	m := minheap.MinHeap{}
	for i := 20; i > 0; i-- {
		m.Insert(i)
		fmt.Println(m)
	}

	for i := 0; i < 10; i++ {
		e := m.Extract()
		fmt.Println("Extracted:", e)
		fmt.Println(m)
	}
}
