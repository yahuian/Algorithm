package _0020

import (
	"Algorithm/DataStructure"
)

func isValid(s string) bool {
	symbolMap := map[string]string{")": "(", "}": "{", "]": "["}
	aStack := DataStructure.Stack{}

	for _, char := range s {
		if _, ok := symbolMap[string(char)]; ok == false {
			aStack.Push(string(char))
		} else if len(aStack) == 0 {
			return false
		} else {
			value, _ := aStack.Pop()
			if symbolMap[string(char)] == value.(string) {
				continue
			} else {
				return false
			}
		}
	}
	if len(aStack) == 0 {
		return true
	} else {
		return false
	}
}
