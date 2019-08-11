package DataStructure

// 单链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 数组转换为单向链表
func ArrayToList(s []int) *ListNode {
	if len(s) == 0 {
		return nil
	}
	head := &ListNode{s[0], nil}
	p := head
	for i := 1; i < len(s); i++ {
		temp := &ListNode{s[i], nil}
		p.Next = temp
		p = p.Next
	}
	return head
}

// 数组转换为有环的单向链表
// pos=-1表示没有环
// pos=0 表示尾节点和首节点相接
// pos=1 表示尾节点和第二个节点相接
func ArrayToCycleList(s []int, pos int) *ListNode {
	// 本函数没有对pos的取值做校验
	if len(s) == 0 {
		return nil
	}
	head := &ListNode{s[0], nil}
	p := head
	for i := 1; i < len(s); i++ {
		temp := &ListNode{s[i], nil}
		p.Next = temp
		p = p.Next
	}
	if pos == -1 {
		return head
	}
	tmp := head
	for i := 0; i < pos; i++ {
		tmp = tmp.Next
	}
	p.Next = tmp
	return head
}
