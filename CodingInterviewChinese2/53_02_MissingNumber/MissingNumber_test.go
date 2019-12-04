package _53_02

import "testing"

func TestMissingNumber(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "缺失的是第一个数字",
			args: args{arr: []int{1, 2, 3, 4, 5}},
			want: 0,
		},
		{
			name: "缺失的是最后一个数字",
			args: args{arr: []int{0, 1, 2, 3, 4}},
			want: 5,
		},
		{
			name: "缺失的是中间某个数字",
			args: args{arr: []int{0, 1, 2, 4, 5}},
			want: 3,
		},
		{
			name: "数组中只有一个数字，缺失的是第一个数字",
			args: args{arr: []int{1}},
			want: 0,
		},
		{
			name: "数组中只有一个数字，缺失的是最后一个数字",
			args: args{arr: []int{0}},
			want: 1,
		},
		{
			name: "鲁棒性",
			args: args{arr: nil},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MissingNumber(tt.args.arr); got != tt.want {
				t.Errorf("MissingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
