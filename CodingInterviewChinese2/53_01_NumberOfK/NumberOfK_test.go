package _53_01

import "testing"

func TestGetNumberOfK(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "查找的数组出现在数组的开头",
			args: args{
				arr: []int{3, 3, 3, 3, 4, 5},
				k:   3,
			},
			want: 4,
		},
		{
			name: "查找的数组出现在数组的结尾",
			args: args{
				arr: []int{1, 2, 3, 3, 3, 3},
				k:   3,
			},
			want: 4,
		},
		{
			name: "查找的数字不存在",
			args: args{
				arr: []int{1, 3, 3, 3, 3, 4, 5},
				k:   2,
			},
			want: 0,
		},
		{
			name: "查找的数字比第一个数字还小，不存在",
			args: args{
				arr: []int{1, 3, 3, 3, 3, 4, 5},
				k:   0,
			},
			want: 0,
		},
		{
			name: "查找的数字比最后一个数字还大，不存在",
			args: args{
				arr: []int{3, 3, 3, 3, 4, 5},
				k:   6,
			},
			want: 0,
		},
		{
			name: "数组中的数字从头到尾都是查找的数字",
			args: args{
				arr: []int{3, 3, 3, 3, 3, 3},
				k:   3,
			},
			want: 6,
		},
		{
			name: "数组中的数字从头到尾只有一个重复的数字，不是查找的数字",
			args: args{
				arr: []int{3, 3, 3, 3, 3, 3},
				k:   4,
			},
			want: 0,
		},
		{
			name: "数组中只有一个数字，是查找的数字",
			args: args{
				arr: []int{3},
				k:   3,
			},
			want: 1,
		},
		{
			name: "数组中只有一个数字，不是查找的数字",
			args: args{
				arr: []int{3},
				k:   4,
			},
			want: 0,
		},
		{
			name: "鲁棒性测试，数组空指针",
			args: args{
				arr: nil,
				k:   0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNumberOfK(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("GetNumberOfK() = %v, want %v", got, tt.want)
			}
		})
	}
}
