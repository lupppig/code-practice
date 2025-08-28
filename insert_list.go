package main

import "fmt"

func (ll *SinglyLinkedList) insertNodeAtHead(data int32) {
	newNode := &SinglyLinkedListNode{data: data}
	if (ll.head) == nil {
		ll.head = newNode
	} else {
		newNode.next = ll.head
		ll.head = newNode
	}
}

func printLinkedList(llist *SinglyLinkedListNode) {
	for llist != nil {
		fmt.Printf("%d->", llist.data)
		llist = llist.next
	}

	fmt.Println()
}
