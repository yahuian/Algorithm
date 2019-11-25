package _0148

import (
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	ptr := head
	var temp []int
	for ptr != nil {
		temp = append(temp, ptr.Val)
		ptr = ptr.Next
	}
	sort.Ints(temp)
	newHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	p := newHead
	for i := 0; i < len(temp); i++ {
		node := &ListNode{
			Val:  temp[i],
			Next: nil,
		}
		p.Next = node
		p = p.Next
	}
	return newHead.Next
}
