package DataStructure

import "errors"

type Queue []interface{}

func (q *Queue) Push(value interface{}) {
	*q = append(*q, value)
}

func (q *Queue) Pop() (interface{}, error) {
	if len(*q) == 0 {
		return nil, errors.New("out of index,the queue is empty")
	}
	value := (*q)[0]
	*q = (*q)[1:len(*q)]
	return value, nil
}

func (q *Queue) Peek() (interface{}, error) {
	if len(*q) == 0 {
		return nil, errors.New("out of index,the queue is empty")
	}
	value := (*q)[0]
	return value, nil
}

func (q *Queue) Empty() bool {
	if len(*q) == 0 {
		return true
	}
	return false
}

func (q *Queue) Len() int {
	return len(*q)
}
