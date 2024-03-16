package examples

import (
	linkedlist "data-structures/types/linkedList"
)

func LinkedListExample() {
	l := linkedlist.LinkedList{}
	data := []int{1, 22, 53, 4, 15, 6, 76, 81, 9, 10}

	for _, val := range data {
		l.Prepend(&linkedlist.Node{Data: val})
	}
	l.PrintList()

	l.DeleteWithValue(4)
	l.DeleteWithValue(4)
	l.DeleteWithValue(4)
	l.DeleteWithValue(4)
	l.DeleteWithValue(4)
	l.PrintList()

	l.DeleteWithValue(53)
	l.PrintList()
}
