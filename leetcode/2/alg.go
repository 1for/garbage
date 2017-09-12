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
	var val int
	var header *ListNode
	
	t1 := l1
	t2 := l2
	carry := false

	//处理头Node
	if t1 != nil && t2 != nil {
		val = t1.Val + t2.Val
		header = NewListNode(val % 10)

		if val >= 10 {
			carry = true
		}
	}
	
	tmp := header
	t1 = t1.Next
	t2 = t2.Next

	//处理后续Node
	for t1 != nil || t2 != nil || carry {
		if carry == true {
			val = 1
		}else{
			val = 0
		}

		if( t1 != nil) {
			val += t1.Val
			t1 = t1.Next
		}

		if( t2 != nil) {
			val += t2.Val
			t2 = t2.Next
		}

		tmp.Next = NewListNode(val % 10);
		tmp = tmp.Next

		if val >= 10 {
			carry = true
		}else{
			carry = false;
		}
	}

	return header
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
