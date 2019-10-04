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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
