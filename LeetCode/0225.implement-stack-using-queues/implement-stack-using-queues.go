package _0225

import "Algorithm/DataStructure"

// 用两个队列实现栈
type MyStack struct {
	q1 DataStructure.Queue
	q2 DataStructure.Queue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	myStack := MyStack{}
	return myStack
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.q1.Push(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	for this.q1.Len() != 1 {
		v, _ := this.q1.Pop()
		this.q2.Push(v)
	}
	value, _ := this.q1.Pop()
	this.q1, this.q2 = this.q2, this.q1
	return value.(int)
}

/** Get the top element. */
func (this *MyStack) Top() int {
	for this.q1.Len() != 1 {
		v, _ := this.q1.Pop()
		this.q2.Push(v)
	}
	value, _ := this.q1.Pop()
	this.q2.Push(value)
	this.q1, this.q2 = this.q2, this.q1
	return value.(int)
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if this.q1.Len() == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
