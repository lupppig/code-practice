package main

import "fmt"

func reversePrint(llist *SinglyLinkedListNode) {
	// Write your code here
	if llist == nil {
		return
	}
	reversePrint(llist.next)
	fmt.Println(llist.data)
}
