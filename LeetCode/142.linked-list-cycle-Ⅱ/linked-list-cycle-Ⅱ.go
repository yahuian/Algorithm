package _142

import "Algorithm/DataStructure"

type ListNode = DataStructure.ListNode

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// slow,fast同时同地出发
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next      // 走一步
		fast = fast.Next.Next // 走二步

		if slow == fast {
			// 有环
			for slow != head {
				slow, head = slow.Next, head.Next
			}
			return slow // 入环点
		}
	}

	// fast为nil，链表到头了，无环
	return nil
}

// https://leetcode-cn.com/problems/linked-list-cycle-ii/solution/huan-xing-lian-biao-ii-by-leetcode/
// 使用快慢指针确定链表有环后，怎么找到入环点，这个题解讲的挺清楚
