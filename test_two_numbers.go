package main

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	node := Node{}
	Ln := &ListNode{Val: 10}
	Ln1 := &ListNode{Val: 20}
	Ln2 := &ListNode{Val: 30}
	Ln3 := &ListNode{Val: 40}

	node.Prepend(Ln)
	node.Prepend(Ln1)
	node.Prepend(Ln2)
	node.Prepend(Ln3)

}
