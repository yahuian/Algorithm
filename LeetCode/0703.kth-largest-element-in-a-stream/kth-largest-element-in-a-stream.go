package _0703

import "container/heap"

type KthLargest struct {
	k int
	h minHeap
}

func Constructor(k int, nums []int) KthLargest {
	// 初始化最小堆
	h := minHeap(nums)
	heap.Init(&h)

	// 维护一个长度为K的最小堆
	for len(h) > k {
		heap.Pop(&h)
	}

	return KthLargest{
		k: k,
		h: h,
	}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(&this.h, val)
	if len(this.h) > this.k {
		heap.Pop(&this.h)
	}
	return this.h[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

// 最小堆接口的实现,官方代码在Go SDK/src/container/heap/example_intheap_test.go
// 即实现Len(),Less(),Swap(),Push(),Pop()这五个方法
type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	result := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return result
}
