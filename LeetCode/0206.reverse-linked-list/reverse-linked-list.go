package _0206

import "Algorithm/DataStructure"

type ListNode = DataStructure.ListNode

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}

// 递归
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}
