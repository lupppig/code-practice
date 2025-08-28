package main

func compare_lists(list1 *SinglyLinkedListNode, list2 *SinglyLinkedListNode) int32 {
	if list2 == nil || list1 == nil {
		return 0
	}

	ll1 := list1
	ll2 := list2

	for ll1 != nil && ll2 != nil {
		if ll1.data != ll2.data {
			return 0
		}
		ll1 = ll1.next
		ll2 = ll2.next
		if ll1 == nil && ll2 == nil {
			return 1
		}
	}

	return 0
}
