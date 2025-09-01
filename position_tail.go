package main

func getNode(llist *SinglyLinkedListNode, positionFromTail int32) int32 {
	// Write your code here
	var sizeList int32
	m := llist
	for m != nil {
		sizeList++
		m = m.next
	}

	for i := sizeList - 1; llist.next != nil && i > 0; i-- {
		if i == positionFromTail {
			break
		}
		llist = llist.next
	}
	return llist.data
}
