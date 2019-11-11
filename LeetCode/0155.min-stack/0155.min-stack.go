package _0155

type MinStack struct {
	s1 []int
	s2 []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.s1 = append(this.s1, x)

	// 长度为0的切片不一定==nil
	// 等于nil的切片长度一定为0
	length := len(this.s2)
	if length == 0 {
		this.s2 = append(this.s2, x)
	} else if x <= this.s2[length-1] {
		this.s2 = append(this.s2, x)
	}
}

func (this *MinStack) Pop() {
	s1TopItem := this.s1[len(this.s1)-1]
	s2TopItem := this.s2[len(this.s2)-1]
	if s1TopItem == s2TopItem {
		this.s2 = this.s2[:len(this.s2)-1]
	}

	this.s1 = this.s1[:len(this.s1)-1]
}

func (this *MinStack) Top() int {
	return this.s1[len(this.s1)-1]
}

func (this *MinStack) GetMin() int {
	return this.s2[len(this.s2)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
