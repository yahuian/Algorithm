package _142

import (
	"Algorithm/DataStructure"
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	head1 := DataStructure.ArrayToCycleList([]int{3, 2, 0, -4}, 1)
	head2 := DataStructure.ArrayToCycleList([]int{1, 2}, 0)
	head3 := DataStructure.ArrayToCycleList([]int{1}, -1)

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// 测试用例
		{
			name: "test1",
			args: args{
				head: head1,
			},
			want: head1.Next,
		},
		{
			name: "test2",
			args: args{
				head: head2,
			},
			want: head2,
		},
		{
			name: "test3",
			args: args{
				head: head3,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
