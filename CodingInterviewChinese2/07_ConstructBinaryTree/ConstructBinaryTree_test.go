package _07

import (
	"reflect"
	"testing"
)

func Test_myTest(t *testing.T) {
	type args struct {
		preTree []int
		inTree  []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "普通二叉树",
			args: args{
				preTree: []int{1, 2, 4, 7, 3, 5, 6, 8},
				inTree:  []int{4, 7, 2, 1, 5, 3, 8, 6},
			},
			want: [][]int{{1, 2, 4, 7, 3, 5, 6, 8}, {4, 7, 2, 1, 5, 3, 8, 6}, {7, 4, 2, 5, 8, 6, 3, 1}},
		},
		{
			name: "所有的节点都没有右子节点",
			args: args{
				preTree: []int{1, 2, 3, 4, 5},
				inTree:  []int{5, 4, 3, 2, 1},
			},
			want: [][]int{{1, 2, 3, 4, 5}, {5, 4, 3, 2, 1}, {5, 4, 3, 2, 1}},
		},
		{
			name: "所有的节点都没有左子节点",
			args: args{
				preTree: []int{1, 2, 3, 4, 5},
				inTree:  []int{1, 2, 3, 4, 5},
			},
			want: [][]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {5, 4, 3, 2, 1}},
		},
		{
			name: "树中只右一个节点",
			args: args{
				preTree: []int{1},
				inTree:  []int{1},
			},
			want: [][]int{{1}, {1}, {1}},
		},
		{
			name: "完全二叉树",
			args: args{
				preTree: []int{1, 2, 4, 5, 3, 6, 7},
				inTree:  []int{4, 2, 5, 1, 6, 3, 7},
			},
			want: [][]int{{1, 2, 4, 5, 3, 6, 7}, {4, 2, 5, 1, 6, 3, 7}, {4, 5, 2, 6, 7, 3, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myTest(tt.args.preTree, tt.args.inTree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("myTest() = %v, want %v", got, tt.want)
			}
		})
	}
}
