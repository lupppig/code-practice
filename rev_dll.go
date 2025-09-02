package main

func reverseDD(llist *DoublyLinkedListNode) *DoublyLinkedListNode {
	// Write your code here

	var prevNode *DoublyLinkedListNode
	// var nextNode *DoublyLinkedListNode
	var currNode *DoublyLinkedListNode = llist

	for currNode != nil {

		prevNode = currNode.prev
		currNode.prev = currNode.next
		currNode.next = prevNode
		currNode = currNode.prev
		// works fine tho...
		// nextNode = currNode.next
		// currNode.next = prevNode
		// prevNode = currNode
		// currNode.prev = nextNode
		// currNode = nextNode
	}
	return prevNode.prev
}
