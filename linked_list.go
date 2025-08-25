package main

type node struct {
	value int
	next  *node
}

type list struct {
	head *node
}

func (l *list) insert(n *node) {
	if l.head == nil {
		l.head = n
	} else {
		curr := l.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = n
	}
}
