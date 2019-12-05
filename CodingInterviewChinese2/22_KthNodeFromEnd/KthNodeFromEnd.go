package _22

// 面试题22：链表中倒数第k个结点
// 题目：输入一个链表，输出该链表中倒数第k个结点。为了符合大多数人的习惯，
// 本题从1开始计数，即链表的尾结点是倒数第1个结点。例如一个链表有6个结点，
// 从头结点开始它们的值依次是1、2、3、4、5、6。这个链表的倒数第3个结点是
// 值为4的结点。

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针
// 注意鲁棒性：
// 1. 链表为nil时
// 2. k小于等于0时（本题从1开始计数）
// 3. 链表长度不够k位时
func findK(head *ListNode, k int) int {
	if head == nil {
		return -1
	}
	if k <= 0 {
		return -1
	}
	f, l := head, head
	// 让快指针先走k-1步
	for i := 0; i < k-1; i++ {
		if f.Next != nil {
			f = f.Next
		} else {
			return -1 // 链表长度不够k位
		}
	}
	// 然后，慢指针和快指针一起走，直到快指针为nil，此时慢指针刚好位于链表倒数第k位
	for f.Next != nil {
		f = f.Next
		l = l.Next
	}
	return l.Val
}
