package main

func deleteNode(llist *SinglyLinkedListNode, position int32) *SinglyLinkedListNode {
	// Write your code here
	var newNode *SinglyLinkedListNode = new(SinglyLinkedListNode)
	if position == 0 {
		return llist.next
	} else {
		newNode = llist
		for i := 1; i < int(position) && newNode != nil; i++ {
			newNode = newNode.next
		}

		if newNode != nil && newNode.next != nil {
			newNode.next = newNode.next.next
		}
	}
	return llist
}
