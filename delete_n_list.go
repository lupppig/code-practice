package main

func deleteNode(llist *SinglyLinkedListNode, position int32) *SinglyLinkedListNode {
	// Write your code here
	var newNode *SinglyLinkedListNode = new(SinglyLinkedListNode)
	if llist == nil {
		return nil
	}
	if position == 0 {
		newNode = llist
		llist = newNode.next
		newNode = nil
	} else {
		newNode = llist
		for i := 1; i < int(position); i++ {
			newNode = newNode.next
		}

		temp := newNode.next
		if temp == nil {
			return nil
		} else {
			newNode.next = temp.next
			temp = nil
		}
	}
	return llist
}
