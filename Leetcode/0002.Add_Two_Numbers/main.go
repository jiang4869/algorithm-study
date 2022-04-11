package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var cur, head *ListNode
	head = &ListNode{}
	cur = head
	carry := 0
	for l1 != nil && l2 != nil {
		num := l1.Val + l2.Val + carry
		carry = num / 10
		num = num % 10
		l1 = l1.Next
		l2 = l2.Next
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
	}
	for l1 != nil {
		num := l1.Val + carry
		carry = num / 10
		num = num % 10
		l1 = l1.Next
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
	}
	for l2 != nil {
		num := l2.Val + carry
		carry = num / 10
		num = num % 10
		l2 = l2.Next
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
	}
	if carry != 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return head.Next
}
