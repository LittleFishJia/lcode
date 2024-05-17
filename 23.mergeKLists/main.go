package main

import (
	"container/heap"
	"errors"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n := &ListNode{Val: 0}
	cur := n
	cur.Next = &ListNode{Val: 2}
	cur.Next = &ListNode{Val: 5}

	printNode(mergeKLists([]*ListNode{n}))
}

func printNode(l *ListNode) {
	cur := l
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	minMap := make(map[int]int)
	for i := 0; i < length; i++ {
		if lists[i] != nil {
			minMap[i] = lists[i].Val
		}
	}

	index, minVal, err := minNums(minMap)
	if err != nil {
		return nil
	}
	delete(minMap, index)

	returnNode := &ListNode{Val: minVal}
	head := returnNode
	for len(minMap) > 0 || lists[index].Next != nil {
		if lists[index].Next != nil {
			minMap[index] = lists[index].Next.Val
			lists[index] = lists[index].Next
		}

		index, minVal, err = minNums(minMap)
		if err != nil {
			break
		}
		delete(minMap, index)

		returnNode.Next = &ListNode{Val: minVal}
		returnNode = returnNode.Next
	}

	return head
}

func minNums(a map[int]int) (int, int, error) {
	if len(a) == 0 {
		return 0, 0, errors.New("err")
	}

	minVal := 0
	index := 0
	first := true
	for i, j := range a {
		if first {
			minVal = j
			index = i
			first = false
		}
		if j < minVal {
			minVal = j
			index = i
		}

	}
	return index, minVal, nil
}

func mergeKList(lists []*ListNode) *ListNode {
	h := hp{}
	for _, head := range lists {
		if head != nil {
			h = append(h, head)
		}
	}
	heap.Init(&h) // 堆化

	dummy := &ListNode{} // 哨兵节点，作为合并后链表头节点的前一个节点
	cur := dummy
	for len(h) > 0 { // 循环直到堆为空
		node := heap.Pop(&h).(*ListNode) // 剩余节点中的最小节点
		if node.Next != nil {            // 下一个节点不为空
			heap.Push(&h, node.Next) // 下一个节点有可能是最小节点，入堆
		}
		cur.Next = node // 合并到新链表中
		cur = cur.Next  // 准备合并下一个节点
	}
	return dummy.Next // 哨兵节点的下一个节点就是新链表的头节点
}

type hp []*ListNode

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].Val < h[j].Val } // 最小堆
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(*ListNode)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
