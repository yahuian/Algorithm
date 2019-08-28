package _0232

import (
	"Algorithm/DataStructure"
)

// 用两个栈来实现队列
type MyQueue struct {
	in  DataStructure.Stack // 压入栈
	out DataStructure.Stack // 弹出栈
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	myQueue := MyQueue{}
	return myQueue
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.in.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	// 因为题目假设所有的操作都是有效的，所以此处不错异常处理

	// 弹出栈为空时，将压入栈中的所有值都弹出到弹出栈保存
	if this.out.Len() == 0 {
		for this.in.Len() != 0 {
			value, _ := this.in.Pop()
			this.out.Push(value)
		}
		v, _ := this.out.Pop()
		return v.(int)
	}
	// 弹出栈非空时，直接从弹出栈中弹出值
	v, _ := this.out.Pop()
	return v.(int)
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.out.Len() == 0 {
		for this.in.Len() != 0 {
			value, _ := this.in.Pop()
			this.out.Push(value)
		}
		v, _ := this.out.Top()
		return v.(int)
	}
	v, _ := this.out.Top()
	return v.(int)
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if this.out.Len() == 0 && this.in.Len() == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
