package _0215

import "container/heap"

func findKthLargest(nums []int, k int) int {
	h := minHeap(nums)
	heap.Init(&h)
	for len(h) > k {
		heap.Pop(&h)
	}
	return heap.Pop(&h).(int)
}

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
