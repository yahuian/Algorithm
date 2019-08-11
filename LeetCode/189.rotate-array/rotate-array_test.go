package _189

import (
	"fmt"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		// 测试用例
		{
			name: "test1",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6},
				k:    3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
			// 输出结果
			fmt.Println(tt.args.nums)
		})
	}
}
