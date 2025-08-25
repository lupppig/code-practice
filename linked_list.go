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

func main() {
	n := &node{value: 1}
	lst := list{}
	n1 := &node{value: 2}
	n2 := &node{value: 3}
	n3 := &node{value: 4}
	n4 := &node{value: 5}
	lst.insert(n)

	lst.insert(n1)
	lst.insert(n2)
	lst.insert(n3)
	lst.insert(n4)

	for n != nil {
		println(n.value)
		n = n.next
	}
}
