package main

import "fmt"

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

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
		fmt.Println(llist.data)
		llist = llist.next
	}
}
