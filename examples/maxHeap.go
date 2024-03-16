package examples

import (
	maxheap "data-structures/types/maxHeap"
	"fmt"
)

func MaxHeapExample() {
	m := maxheap.MaxHeap{}
	for i := 0; i < 20; i++ {
		m.Insert(i)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		e := m.Extract()
		fmt.Println("Extracted:", e)
		fmt.Println(m)
	}
}
