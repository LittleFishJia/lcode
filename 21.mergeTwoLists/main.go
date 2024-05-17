package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n := &ListNode{Val: 0}
	addNode(n, 1)
	addNode(n, 2)
	addNode(n, 4)
	addNode(n, 6)
	a := &ListNode{Val: 3}
	fmt.Println("test")
	b := mergeTwoLists(n, a)

	printNode(b)
}

func addNode(l *ListNode, val int) *ListNode {
	cur := l
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &ListNode{Val: val}
	return l
}

func printNode(l *ListNode) {
	current := l
	for current != nil {
		println(current.Val)
		current = current.Next
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				cur.Next = &ListNode{Val: list1.Val}
				list1 = list1.Next
			} else {
				cur.Next = &ListNode{Val: list2.Val}
				list2 = list2.Next
			}
			cur = cur.Next
			continue
		}
		if list1 != nil {
			cur.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
			cur = cur.Next
		}
		if list2 != nil {
			cur.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
			cur = cur.Next
		}
	}
	return head.Next
}
