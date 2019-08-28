package DataStructure

import "errors"

type Stack []interface{}

func (s Stack) Len() int {
	return len(s)
}

func (s *Stack) Push(value interface{}) {
	*s = append(*s, value)
}

func (s *Stack) Pop() (interface{}, error) {
	if len(*s) == 0 {
		return nil, errors.New("out of index,len is 0")
	}
	value := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return value, nil
}

func (s *Stack) Top() (interface{}, error) {
	if len(*s) == 0 {
		return nil, errors.New("out of index,len is 0")
	}
	return (*s)[len(*s)-1], nil
}
