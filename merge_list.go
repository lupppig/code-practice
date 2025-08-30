package main

func mergeLists(headA *SinglyLinkedListNode, headB *SinglyLinkedListNode) *SinglyLinkedListNode {
	var newList = &SinglyLinkedListNode{}
	tail := newList

	if headA == nil && headB == nil {
		return newList
	} else if headA == nil {
		return headB
	} else if headB == nil {
		return headA
	}

	for headA != nil && headB != nil {
		if headA.data < headB.data {
			tail.next = headA
			headA = headA.next
		} else {
			tail.next = headB
			headB = headB.next
		}
		tail = tail.next
	}

	if headA != nil {
		tail.next = headA
	} else if headB != nil {
		tail.next = headB
	}
	return newList.next
}
