package _33

import "testing"

func TestSequenceOfBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "完全的二叉搜索树",
			args: args{nums: []int{4, 8, 6, 12, 16, 14, 10}},
			want: true,
		},
		{
			name: "普通的二叉搜索树",
			args: args{nums: []int{4, 6, 7, 5}},
			want: true,
		},
		{
			name: "右结点全为空",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: true,
		},
		{
			name: "左结点全为空",
			args: args{nums: []int{5, 4, 3, 2, 1}},
			want: true,
		},
		{
			name: "只有一个结点",
			args: args{nums: []int{4}},
			want: true,
		},
		{
			name: "非二叉搜索树1",
			args: args{nums: []int{7, 4, 6, 5}},
			want: false,
		},
		{
			name: "非二叉搜索树2",
			args: args{nums: []int{4, 6, 12, 8, 16, 14, 10}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SequenceOfBST(tt.args.nums); got != tt.want {
				t.Errorf("SquenceOfBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
