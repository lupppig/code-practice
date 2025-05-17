package main

import (
	"fmt"
	"runtime"
)

/*
*
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order, and each of their nodes contains a single digit.
Add the two numbers and return the sum as a linked list.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Head *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	runtime.GC()
	dummy := &ListNode{}
	addTwoAux(l1, l2, dummy, 0)
	return dummy.Next
}

func addTwoAux(l1 *ListNode, l2 *ListNode, newNode *ListNode, carry int) {
	if l1 == nil && l2 == nil && carry == 0 {
		return
	}
	val1, val2 := 0, 0

	var next1, next2 *ListNode
	if l1 != nil {
		val1 = l1.Val
		next1 = l1.Next
	}
	if l2 != nil {
		val2 = l2.Val
		next2 = l2.Next
	}
	total := val1 + val2 + carry
	carry = int(total / 10)
	new_digit := total % 10

	newNode.Next = &ListNode{Val: new_digit}
	addTwoAux(next1, next2, newNode.Next, carry)
}

func (l *Node) Prepend(val *ListNode) {
	second := l.Head
	l.Head = val
	l.Head.Next = second
}

func (l *ListNode) PrintNode() {
	for l != nil {
		if l.Next != nil {
			fmt.Printf("(%d)---> ", l.Val)
		} else {
			fmt.Printf("(%d)", l.Val)
		}
		l = l.Next
	}
	fmt.Println()
}

func main() {
	node2 := Node{}
	Lnn := &ListNode{Val: 9}
	Lnn1 := &ListNode{Val: 9}
	Lnn2 := &ListNode{Val: 9}
	Lnn3 := &ListNode{Val: 9}
	Lnn4 := &ListNode{Val: 9}
	Lnn5 := &ListNode{Val: 9}
	Lnn6 := &ListNode{Val: 9}

	node2.Prepend(Lnn)
	node2.Prepend(Lnn1)
	node2.Prepend(Lnn2)
	node2.Prepend(Lnn3)
	node2.Prepend(Lnn4)
	node2.Prepend(Lnn5)
	node2.Prepend(Lnn6)

	node := Node{}
	Ln := &ListNode{Val: 9}
	Ln1 := &ListNode{Val: 9}
	Ln2 := &ListNode{Val: 9}
	Ln3 := &ListNode{Val: 9}

	node.Prepend(Ln)
	node.Prepend(Ln1)
	node.Prepend(Ln2)
	node.Prepend(Ln3)

	a := addTwoNumbers(node2.Head, node.Head)
	a.PrintNode()
}
