package _32_02

import (
	. "Algorithm/DataStructure"
	"reflect"
	"testing"
)

func TestPrintTreesInLines(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "普通的二叉树",
			args: args{root: SliceToBinaryTree([]int{8, 6, 10, 5, 7, 9, 11})},
			want: [][]int{{8}, {6, 10}, {5, 7, 9, 11}},
		},
		{
			name: "右结点全为空",
			args: args{root: SliceToBinaryTree([]int{5, 4, NULL, 3, NULL, 2, NULL, 1, NULL})},
			want: [][]int{{5}, {4}, {3}, {2}, {1}},
		},
		{
			name: "左结点全为空",
			args: args{root: SliceToBinaryTree([]int{1, NULL, 2, NULL, 3, NULL, 4, NULL, 5})},
			want: [][]int{{1}, {2}, {3}, {4}, {5}},
		},
		{
			name: "树中只有一个结点",
			args: args{root: SliceToBinaryTree([]int{1})},
			want: [][]int{{1}},
		}, {
			name: "树中没有结点",
			args: args{root: SliceToBinaryTree([]int{})},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintTreesInLines(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrintTreesInLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
