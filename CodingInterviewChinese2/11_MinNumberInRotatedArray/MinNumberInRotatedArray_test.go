package _11

import "testing"

func TestMinNumberInRotatedArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "典型输入，单调升序的数组的一个旋转",
			args: args{nums: []int{3, 4, 5, 1, 2}},
			want: 1,
		},
		{
			name: "有重复数字，并且重复的数字刚好的最小的数字",
			args: args{nums: []int{3, 4, 5, 1, 1, 2}},
			want: 1,
		},
		{
			name: "有重复数字，但重复的数字不是第一个数字和最后一个数字",
			args: args{nums: []int{3, 4, 5, 1, 2, 2}},
			want: 1,
		},
		{
			name: "有重复的数字，并且重复的数字刚好是第一个数字和最后一个数字",
			args: args{nums: []int{1, 0, 1, 1, 1}},
			want: 0,
		},
		{
			name: "单调升序数组，旋转0个元素，也就是单调升序数组本身",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 1,
		},
		{
			name: "数组中只有一个数字",
			args: args{nums: []int{2}},
			want: 2,
		},
		{
			name: "nil",
			args: args{nums: []int{}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinNumberInRotatedArray(tt.args.nums); got != tt.want {
				t.Errorf("MinNumberInRotatedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
