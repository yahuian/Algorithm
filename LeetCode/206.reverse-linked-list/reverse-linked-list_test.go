package _206

import (
	"Algorithm/DataStructure"
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	head1 := DataStructure.ArrayToList([]int{1, 2, 3, 4, 5})
	head2 := DataStructure.ArrayToList([]int{5, 4, 3, 2, 1})

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
			want: head2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
