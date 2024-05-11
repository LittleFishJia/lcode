package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main()  {
	n := &ListNode{Val: 0}
	addNode(n, 1)
	addNode(n, 2)
	addNode(n, 3)
	addNode(n, 4)
	printNode(n)
	fmt.Println("test")
	removeNthFromEnd(n,2)

	printNode(n)
}

func addNode(l *ListNode, val int) *ListNode {
	cur := l
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &ListNode{Val: val}
	return l
}

func printNode(l *ListNode)  {
	current := l
	for current != nil {
		current = current.Next
	}
}

func removeNthFromEnd(l *ListNode, n int) *ListNode {
	count := getLength(l)
	cur := &ListNode{Val: 0, Next: l}
	head := cur
	for i:=0 ;i < count - n ;i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return head
}

func getLength(l *ListNode) int {
	cur := l
	length := 0
	for cur != nil {
		length++
		cur=cur.Next
	}
	return length
}