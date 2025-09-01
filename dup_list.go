package main

func removeDuplicates(llist *SinglyLinkedListNode) *SinglyLinkedListNode {
	if llist == nil {
		return nil
	}

	var futureNode *SinglyLinkedListNode = llist.next
	var prevNode *SinglyLinkedListNode = llist

	for futureNode != nil {
		if futureNode.data == prevNode.data {
			prevNode.next = futureNode.next
		} else {
			prevNode = futureNode
		}
		futureNode = futureNode.next
	}
	return llist
}
