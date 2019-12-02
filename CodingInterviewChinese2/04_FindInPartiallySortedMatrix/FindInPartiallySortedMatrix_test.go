package _04

import "testing"

func TestFindInPartiallySortedMatrix(t *testing.T) {
	type args struct {
		nums [][]int
		key  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "要查找的数在数组中",
			args: args{
				nums: [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}},
				key:  7,
			},
			want: true,
		},
		{
			name: "要查找的数不在数组中",
			args: args{
				nums: [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}},
				key:  5,
			},
			want: false,
		},
		{
			name: "要查找的数是数组中最小的数",
			args: args{
				nums: [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}},
				key:  1,
			},
			want: true,
		},
		{
			name: "要查找的数是数组中最大的数",
			args: args{
				nums: [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}},
				key:  15,
			},
			want: true,
		},
		{
			name: "要查找的数比数组中最小的数字还小",
			args: args{
				nums: [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}},
				key:  0,
			},
			want: false,
		},
		{
			name: "查找的数比数组中最大的数字还大",
			args: args{
				nums: [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}},
				key:  16,
			},
			want: false,
		},
		{
			name: "测试鲁棒性",
			args: args{
				nums: [][]int{},
				key:  16,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindInPartiallySortedMatrix(tt.args.nums, tt.args.key); got != tt.want {
				t.Errorf("FindInPartiallySortedMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
