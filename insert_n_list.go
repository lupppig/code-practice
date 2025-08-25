package main

/*
 * Complete the 'insertNodeAtPosition' function below.
 *
 * The function is expected to return an INTEGER_SINGLY_LINKED_LIST.
 * The function accepts following parameters:
 *  1. INTEGER_SINGLY_LINKED_LIST llist
 *  2. INTEGER data
 *  3. INTEGER position
 */

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */

func (singlyLinkedList *SinglyLinkedList) insertNodeIntoSinglyLinkedList(nodeData int32) {
	node := &SinglyLinkedListNode{
		next: nil,
		data: nodeData,
	}

	if singlyLinkedList.head == nil {
		singlyLinkedList.head = node
	} else {
		singlyLinkedList.tail.next = node
	}

	singlyLinkedList.tail = node
}

func insertNodeAtPosition(llist *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {
	var newNode *SinglyLinkedListNode = &SinglyLinkedListNode{data: data}
	newList := SinglyLinkedList{}
	if llist == nil && position == 0 {
		newList.insertNodeIntoSinglyLinkedList(data)
	} else if position == 0 {
		newNode.next = llist
		llist = newNode
	} else {
		curr := llist
		for i := 1; i < int(position) && curr != nil; i++ {
			curr = curr.next
		}
		temp := curr.next
		curr.next = newNode
		newNode.next = temp
	}
	return llist
}
