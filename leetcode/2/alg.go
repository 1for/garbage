package main

import "fmt"
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode,  l2 *ListNode) *ListNode {
	var val, x, y int
	var header *ListNode = NewListNode(0)
	
	t1 := l1
	t2 := l2
	carry := 0
	tmp := header

	for t1 != nil || t2 != nil {
		x = 0
		y = 0
		
		if t1 != nil {
			x = t1.Val
			t1 = t1.Next
		}

		if t2 != nil {
			y = t2.Val
			t2 = t2.Next
		}

		val = carry + x + y; 
		tmp.Next = NewListNode(val % 10)
		tmp = tmp.Next

		carry = val / 10
	}

	if carry > 0 {
		tmp.Next = NewListNode(carry)
	}

	return header.Next
}

func NewListNode(val int) *ListNode{
	return &ListNode{Val: val}
}

func PrintList(l *ListNode){
	t := l
	fmt.Printf("(")
	for t != nil {
		fmt.Printf("%d ", t.Val)
		t = t.Next
	}
	fmt.Printf(")\n")
}
