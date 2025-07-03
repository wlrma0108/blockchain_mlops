package main

type ListNode struct {
	Value int
	Next  *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func (l *LinkedList) InsertAtFromt(v int) {
	node := &ListNode{Value: v, Next: l.Head}
	l.Head = node
}

func (l *LinkedList) DeleteValue(v int) bool {
	if l.Head == nil {
		return false
	}

	if l.Head.Value == v {
		l.Head = l.Head.Next
		return true
	}

	current := l.Head
	for current.Next != nil {
		if current.Next.Value == v {
			current.Next = current.Next.Next
			return true
		}
		current = current.Next
	}
	return false
}
