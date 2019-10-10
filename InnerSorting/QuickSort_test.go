package InnerSorting

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// 测试用例
		{
			name: "test1",
			args: args{
				s: []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "test2",
			args: args{
				s: []int{},
			},
			want: []int{},
		},
		{
			name: "test3",
			args: args{
				s: []int{2, 1},
			},
			want: []int{1, 2},
		},
		{
			name: "test4",
			args: args{
				s: []int{1, 2, 3, 8, 6, 4},
			},
			want: []int{1, 2, 3, 4, 6, 8},
		},
		{
			name: "test5",
			args: args{
				s: []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586},
			},
			want: []int{-5467984, -784, 0, 0, 42, 59, 74, 238, 905, 959, 7586, 7586, 9845},
		},
		{
			name: "test6",
			args: args{
				s: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "test7",
			args: args{
				s: []int{1, 2, 3, 3, 2, 1},
			},
			want: []int{1, 1, 2, 2, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
