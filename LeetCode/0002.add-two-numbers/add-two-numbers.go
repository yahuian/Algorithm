package _0002

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p, q := l1, l2
	head := &ListNode{0, nil}
	pre := head
	carry := 0 // 进位
	for p != nil || q != nil {
		var x, y int
		if p != nil {
			x = p.Val
		} else {
			x = 0
		}
		if q != nil {
			y = q.Val
		} else {
			y = 0
		}

		sum := x + y + carry
		carry = sum / 10
		pre.Next = &ListNode{sum % 10, nil}
		pre = pre.Next

		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if carry > 0 {
		pre.Next = &ListNode{carry, nil}
	}
	return head.Next
}
