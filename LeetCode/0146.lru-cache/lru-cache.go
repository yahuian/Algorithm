package _0146

import (
	"container/list"
)

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

type LRUCache struct {
	hashMap    map[int]*list.Element
	cap        int
	twoWayList *list.List
}

type pair struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		hashMap:    make(map[int]*list.Element, capacity),
		cap:        capacity,
		twoWayList: new(list.List),
	}
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.hashMap[key]; ok {
		value := node.Value.(*list.Element).Value.(pair).value
		l.twoWayList.MoveToFront(node)
		return value
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	// 检查key是否存在
	if node, ok := l.hashMap[key]; ok {
		l.twoWayList.MoveToFront(node)
		node.Value.(*list.Element).Value = pair{key: key, value: value}
	} else {
		// 容量不够
		if len(l.hashMap) == l.cap {
			// 获取要删除节点的key，即双向链表的尾节点
			index := l.twoWayList.Back().Value.(*list.Element).Value.(pair).key
			delete(l.hashMap, index)
			l.twoWayList.Remove(l.twoWayList.Back())
		}
		node := &list.Element{
			Value: pair{key: key, value: value},
		}
		ptr := l.twoWayList.PushFront(node)
		l.hashMap[key] = ptr
	}
}

// https://github.com/aQuaYi/LeetCode-in-Go/blob/master/Algorithms/0146.lru-cache/lru-cache.go
