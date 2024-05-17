package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func createLinkedList(nums []int) *ListNode {
	if len(nums) <= 0 {
		return nil
	}
	head := &ListNode{}
	current := head
	numsLen := len(nums)
	for index, num := range nums {
		current.Val = num
		if numsLen-1 > index {
			current.Next = &ListNode{}
			current = current.Next
		}
	}
	return head
}

func printList(node *ListNode) {
	if node == nil {
		return
	}
	//current := node
	for node != nil {
		println(node.Val)
		node = node.Next
	}
}

func main() {
	lista := createLinkedList([]int{2, 7, 11, 15, 9})
	listb := createLinkedList([]int{1, 2, 3, 4})
	listc := addTwoNumbers(lista, listb)
	printList(listc)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	newNode := &ListNode{}
	head := newNode
	var addNums = 0
	for l1 != nil || l2 != nil || addNums > 0 {
		if l1 != nil {
			addNums += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			addNums += l2.Val
			l2 = l2.Next
		}
		newNode.Val = addNums % 10
		addNums = addNums / 10
		if l1 != nil || l2 != nil || addNums > 0 {
			newNode.Next = &ListNode{Val: addNums}
			newNode = newNode.Next
		}
	}
	return head
}

//
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	newNode := &ListNode{}
//	head := newNode
//	var addNums = 0
//	for l1 != nil || l2 != nil || addNums > 0 {
//		if l1 != nil {
//			addNums += l1.Val
//			l1 = l1.Next
//		}
//		if l2 != nil {
//			addNums += l2.Val
//			l2 = l2.Next
//		}
//		newNode.Val = addNums % 10
//		addNums = addNums / 10
//		if l1 != nil || l2 != nil || addNums > 0 {
//			newNode.Next = &ListNode{Val: addNums}
//			newNode = newNode.Next
//		}
//	}
//	return head
//}
