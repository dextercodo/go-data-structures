package linkedlist

import "fmt"

type Node struct {
	Data int
	next *Node
}

func (n *Node) Next() *Node {
	return n.next
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) Head() *Node {
	return l.head
}

func (l *LinkedList) Prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l *LinkedList) DeleteWithValue(val int) {
	if l.length == 0 {
		return
	}
	if l.head.Data == val {
		l.head = l.head.next
		l.length--
		return
	}
	previousToDelete := l.head
	for previousToDelete.next != nil {
		if previousToDelete.next.Data == val {
			previousToDelete.next = previousToDelete.next.next
			l.length--
			return
		}
		previousToDelete = previousToDelete.next
	}
}

func (l LinkedList) PrintList() {
	toPrint := l.head
	for l.length != 0 {
		println("next: ", toPrint.Data)
		toPrint = toPrint.Next()
		l.length--
	}
	fmt.Println("------- End of List -------")
}
