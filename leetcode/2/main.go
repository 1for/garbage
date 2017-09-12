package main

func main(){
	var l1 *ListNode
	var l2 *ListNode

	l1 = NewListNode(2)
	l1.Next = NewListNode(4)
	l1.Next.Next = NewListNode(3)

	l2 = NewListNode(5)
	l2.Next = NewListNode(6)
	l2.Next.Next = NewListNode(4)

	result := addTwoNumbers(l1, l2)

	PrintList(l1)
	PrintList(l2)
	PrintList(result)
}
