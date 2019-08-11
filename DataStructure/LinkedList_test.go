package DataStructure

import (
	"reflect"
	"testing"
)

func TestArrayToList(t *testing.T) {
	type args struct {
		s []int
	}
	// 创建单链表
	head := &ListNode{1, nil}
	p1 := &ListNode{2, nil}
	p2 := &ListNode{3, nil}
	head.Next = p1
	p1.Next = p2

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// 测试用例
		{
			name: "test1",
			args: args{
				s: []int{1, 2, 3},
			},
			want: head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayToList(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayToList() = %v, want %v", got, tt.want)
			}
		})
	}
}
