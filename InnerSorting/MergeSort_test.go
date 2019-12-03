package InnerSorting

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "普通情况，且有重复元素",
			args: args{
				arr: []int{25, 34, 45, 32, 78, 12, 34, 64},
			},
			want: []int{12, 25, 32, 34, 34, 45, 64, 78},
		},
		{
			name: "普通情况，且无重复元素",
			args: args{
				arr: []int{25, 3, 45, 32, 78, 12, 34, 64},
			},
			want: []int{3, 12, 25, 32, 34, 45, 64, 78},
		},
		{
			name: "原数组本来有序，正序",
			args: args{
				arr: []int{1, 3, 5, 7, 9, 12, 34, 64},
			},
			want: []int{1, 3, 5, 7, 9, 12, 34, 64},
		},
		{
			name: "原数组本来有序，倒序",
			args: args{
				arr: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "空数组",
			args: args{
				arr: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
