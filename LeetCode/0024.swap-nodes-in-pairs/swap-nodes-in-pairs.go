package _0024

import "Algorithm/DataStructure"

type ListNode = DataStructure.ListNode

func swapPairs(head *ListNode) *ListNode {
	// 输入校验
	if head == nil {
		return nil
	}
	pre := &ListNode{}
	pre.Next = head
	tmp := pre
	for tmp.Next != nil && tmp.Next.Next != nil {
		p1 := tmp.Next
		p2 := tmp.Next.Next

		tmp.Next = p2
		p1.Next = p2.Next
		p2.Next = p1
		tmp = p1
	}
	return pre.Next
}

// https://leetcode-cn.com/problems/swap-nodes-in-pairs/solution/hua-jie-suan-fa-24-liang-liang-jiao-huan-lian-biao/

// 递归法
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var tmp = head.Next
	var after = swapPairs2(tmp.Next)
	head.Next = after
	tmp.Next = head
	return tmp
}
