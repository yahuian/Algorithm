package _32_01

import (
	. "Algorithm/DataStructure"
	"reflect"
	"testing"
)

func TestPrintTreeFromTopToBottom(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "普通的二叉树",
			args: args{root: SliceToBinaryTree([]int{10, 6, 14, 4, 8, 12, 16})},
			want: []int{10, 6, 14, 4, 8, 12, 16},
		},
		{
			name: "右结点全为空",
			args: args{root: SliceToBinaryTree([]int{5, 4, NULL, 3, NULL, 2, NULL, 1, NULL})},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "左结点全为空",
			args: args{root: SliceToBinaryTree([]int{1, NULL, 2, NULL, 3, NULL, 4, NULL, 5})},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "树中只有一个结点",
			args: args{root: SliceToBinaryTree([]int{1})},
			want: []int{1},
		}, {
			name: "树中没有结点",
			args: args{root: SliceToBinaryTree([]int{})},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintTreeFromTopToBottom(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrintTreeFromTopToBottom() = %v, want %v", got, tt.want)
			}
		})
	}
}
