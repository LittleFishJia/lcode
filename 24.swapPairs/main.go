package main

import "fmt"

func main() {
	L := &ListNode{Val: 2}
	addNode(L, 1)
	addNode(L, 4)
	printNode(L)
	swapPairs(L)
	fmt.Println("test")
	printNode(L)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addNode(l *ListNode, val int) {
	cur := l
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &ListNode{Val: val}
}

func printNode(l *ListNode) {
	cur := l
	for cur != nil {
		println(cur.Val)
		cur = cur.Next
	}
}

func getLength(l *ListNode) int {
	cur := l
	length := 0
	for cur != nil {
		length++
		cur = cur.Next
	}
	return length
}
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
}
