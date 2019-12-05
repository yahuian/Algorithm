package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 面试题6：从尾到头打印链表
// 题目：输入一个链表的头结点，从尾到头反过来打印出每个结点的值。

// 方法0：借助栈，实现后进先出

// 方法1： 用递归的方法来写，且不修改原链表的结构
func reverseList(head *ListNode) {
	if head != nil {
		if head.Next != nil {
			reverseList(head.Next)
		}
		fmt.Print(head.Val)
	}
}

func main() {
	p5 := &ListNode{Val: 5, Next: nil}
	p4 := &ListNode{Val: 4, Next: p5}
	p3 := &ListNode{Val: 3, Next: p4}
	p2 := &ListNode{Val: 2, Next: p3}
	p1 := &ListNode{Val: 1, Next: p2}

	reverseList(p1) // 54321
}
