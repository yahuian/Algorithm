package _0002

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 方法1：先转化为数字，再转化回链表
func addTwoNumbersOne(l1 *ListNode, l2 *ListNode) *ListNode {
	s := sum(l1) + sum(l2)
	head := new(ListNode)
	h := head
	for {
		v := s % 10
		h.Val = v
		s /= 10
		if s == 0 {
			return head
		}
		h.Next = new(ListNode)
		h = h.Next
	}
}

func sum(l *ListNode) int {
	var i, s int
	for l != nil {
		s += l.Val * int(math.Pow(float64(10), float64(i)))
		i++
		l = l.Next
	}
	return s
}

// 方法2：
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
