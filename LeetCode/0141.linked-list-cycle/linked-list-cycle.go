package _0141

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	p1 := head
	p2 := head
	for p1 != nil && p2 != nil && p2.Next != nil {
		p1 = p1.Next      // p1每次走一步
		p2 = p2.Next.Next // p2每次走二步
		// 如果p1和p2相遇，则代表有环
		if p1 == p2 {
			return true
		}
	}
	return false
}
