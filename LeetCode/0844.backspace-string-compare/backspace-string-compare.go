package _0844

import "Algorithm/DataStructure"

func backspaceCompare(S string, T string) bool {
	stack1 := DataStructure.Stack{}
	for _, v := range S {
		if string(v) == "#" && stack1.Len() == 0 {
			continue
		}
		if string(v) == "#" {
			stack1.Pop()
			continue
		}
		stack1.Push(string(v))
	}

	stack2 := DataStructure.Stack{}
	for _, v := range T {
		if string(v) == "#" && stack2.Len() == 0 {
			continue
		}
		if string(v) == "#" {
			stack2.Pop()
			continue
		}
		stack2.Push(string(v))
	}

	// 暴力比较两个栈的内容是否相同
	for stack1.Len() != 0 && stack2.Len() != 0 {
		v1, _ := stack1.Pop()
		v2, _ := stack2.Pop()
		if v1 != v2 {
			return false
		}
	}

	if stack1.Len() == 0 && stack2.Len() == 0 {
		return true
	} else {
		return false
	}
}
