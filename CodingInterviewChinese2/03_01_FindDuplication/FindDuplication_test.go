package _03_01

import (
	"reflect"
	"testing"
)

func TestFindDuplication(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "重复的数字是数组中最小的数字",
			args: args{
				nums: []int{2, 1, 3, 1, 4},
			},
			want: []int{1},
		},
		{
			name: "重复的数字是数组中最大的数字",
			args: args{
				nums: []int{2, 4, 3, 1, 4},
			},
			want: []int{4},
		},
		{
			name: "数组中存在多个重复的数字(1)",
			args: args{
				nums: []int{2, 4, 2, 1, 4},
			},
			want: []int{2, 4},
		},
		{
			name: "数组中存在多个重复的数字(2)",
			args: args{
				nums: []int{2, 3, 1, 0, 2, 5, 3},
			},
			want: []int{2, 3},
		},
		{
			name: "没有重复的数字",
			args: args{
				nums: []int{2, 1, 3, 0, 4},
			},
			want: nil,
		},
		{
			name: "无效的输入",
			args: args{
				nums: []int{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDuplication(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindDuplication() = %v, want %v", got, tt.want)
			}
		})
	}
}
