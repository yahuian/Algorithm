package _0061

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// to ring and count the length of list
	n := 1
	temp := head
	for temp.Next != nil {
		n++
		temp = temp.Next
	}
	temp.Next = head

	// new head = n-(k%n)
	// new tail = n-(k%n)-1
	newTail := head
	for i := 0; i < n-(k%n)-1; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next

	// break the ring
	newTail.Next = nil
	return newHead
}
