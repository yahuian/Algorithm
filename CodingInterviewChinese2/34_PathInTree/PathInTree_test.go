package _34

import (
	. "Algorithm/DataStructure"
	"reflect"
	"testing"
)

func TestPathInTree(t *testing.T) {
	type args struct {
		root *TreeNode
		key  int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "有两条路径上的结点和为22",
			args: args{
				root: SliceToBinaryTree([]int{10, 5, 12, 4, 7}),
				key:  22,
			},
			want: [][]int{{10, 5, 7}, {10, 12}},
		},
		{
			name: "没有路径上的结点和为15",
			args: args{
				root: SliceToBinaryTree([]int{10, 5, 12, 4, 7}),
				key:  15,
			},
			want: nil,
		},
		{
			name: "有一条路径上的结点和为15",
			args: args{
				root: SliceToBinaryTree([]int{5, 4, NULL, 3, NULL, 2, NULL, 1, NULL}),
				key:  15,
			},
			want: [][]int{{5, 4, 3, 2, 1}},
		},
		{
			name: "没有路径上的和为16",
			args: args{
				root: SliceToBinaryTree([]int{1, NULL, 2, NULL, 3, NULL, 4, NULL, 5}),
				key:  16,
			},
			want: nil,
		},
		{
			name: "树中只有一个结点",
			args: args{
				root: SliceToBinaryTree([]int{10}),
				key:  10,
			},
			want: [][]int{{10}},
		},
		{
			name: "树中无结点",
			args: args{
				root: SliceToBinaryTree([]int{}),
				key:  22,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PathInTree(tt.args.root, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PathInTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
