package main

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_Case1(t *testing.T) {
	arr1 := []int{2, 4, 3}
	arr2 := []int{5, 6, 4}
	l1 := CreateList(arr1)
	l2 := CreateList(arr2)
	numbers := addTwoNumbers(l1, l2)
	PrintList(numbers)
}

func Test_Case2(t *testing.T) {
	arr1 := []int{0}
	arr2 := []int{0}
	l1 := CreateList(arr1)
	l2 := CreateList(arr2)
	numbers := addTwoNumbers(l1, l2)
	PrintList(numbers)
}

func Test_Case3(t *testing.T) {
	arr1 := []int{9, 9, 9, 9, 9, 9, 9}
	arr2 := []int{9, 9, 9, 9}
	l1 := CreateList(arr1)
	l2 := CreateList(arr2)
	numbers := addTwoNumbers(l1, l2)
	PrintList(numbers)
}

func CreateList(nums []int) *ListNode {
	var cur, head *ListNode
	head = &ListNode{}
	cur = head
	for _, val := range nums {
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}
	return head.Next
}

func PrintList(list *ListNode) {
	var str string
	str += "["
	flag := false
	for list != nil {
		if flag {
			str += ","
		}
		str += strconv.FormatInt(int64(list.Val), 10)

		flag = true
		list = list.Next
	}
	str += "]"
	fmt.Println(str)
}
