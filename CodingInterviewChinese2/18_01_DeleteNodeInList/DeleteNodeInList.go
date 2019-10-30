package main

import (
	"fmt"
)

// 单链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteNodeInList(head, delete *ListNode) *ListNode {
	if head == nil || delete == nil {
		return head
	}
	if delete == head {
		return head.Next
	}
	// 下面的情况中，链表至少有两个节点
	if delete.Next == nil {
		p := head
		for p.Next != delete {
			p = p.Next
		}
		p.Next = nil

	} else {
		delete.Val = delete.Next.Val
		delete.Next = delete.Next.Next
	}
	return head
}

// 输出链表
func show(head *ListNode) {
	if head == nil {
		fmt.Print("nil")
	}
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Printf("\n")
}

func main() {
	test1()
	fmt.Println("--------------")
	test2()
	fmt.Println("--------------")
	test3()
	fmt.Println("--------------")
	test4()
}

func test1() {
	// 链表中有多个节点，删除中间的节点
	p5 := &ListNode{5, nil}
	p4 := &ListNode{4, p5}
	p3 := &ListNode{3, p4}
	p2 := &ListNode{2, p3}
	p1 := &ListNode{1, p2}
	show(p1)
	test1 := DeleteNodeInList(p1, p3)
	show(test1)
}

func test2() {
	// 链表中有多个节点，删除尾节点
	p5 := &ListNode{5, nil}
	p4 := &ListNode{4, p5}
	p3 := &ListNode{3, p4}
	p2 := &ListNode{2, p3}
	p1 := &ListNode{1, p2}
	show(p1)
	test := DeleteNodeInList(p1, p5)
	show(test)
}

func test3() {
	// 链表中有多个节点，删除头节点
	p5 := &ListNode{5, nil}
	p4 := &ListNode{4, p5}
	p3 := &ListNode{3, p4}
	p2 := &ListNode{2, p3}
	p1 := &ListNode{1, p2}
	show(p1)
	test := DeleteNodeInList(p1, p1)
	show(test)
}

func test4() {
	// 链表为空
	show(nil)
	test := DeleteNodeInList(nil, nil)
	show(test)
}
