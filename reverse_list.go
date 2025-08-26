package main

func reverse(llist *SinglyLinkedListNode) *SinglyLinkedListNode {
	if llist == nil {
		return nil
	}
	currNode := llist
	var prevNode *SinglyLinkedListNode
	var nextNode *SinglyLinkedListNode
	for currNode != nil {
		nextNode = currNode.next
		currNode.next = prevNode
		prevNode = currNode
		currNode = nextNode
	}

	llist = prevNode
	return llist
}
