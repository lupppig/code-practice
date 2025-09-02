package main

func sortedInsert(llist *DoublyLinkedListNode, data int32) *DoublyLinkedListNode {
	// Write your code here

	var newLL *DoublyLinkedListNode = &DoublyLinkedListNode{data: data}
	if llist == nil {
		return newLL
	}
	ll := llist

	for ll != nil {
		if newLL.data >= ll.data && ll.next == nil {
			ll.next = newLL
			newLL.prev = ll
			break
		} else if newLL.data <= ll.data && ll.prev == nil {
			newLL.next = ll
			llist = newLL
			ll.prev = newLL
			break
		} else if newLL.data <= ll.data {
			newLL.next = ll
			newLL.prev = ll.prev
			ll.prev.next = newLL
			ll.prev = newLL
			break
		}
		ll = ll.next
	}
	return llist
}
