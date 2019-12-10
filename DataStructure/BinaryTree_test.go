package DataStructure

import (
	"reflect"
	"testing"
)

func Test_myTest(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "传入nil切片",
			args: args{s: nil},
			want: [][]int{nil, nil, nil},
		},
		{
			name: "只有一个根节点",
			args: args{s: []int{1}},
			want: [][]int{{1}, {1}, {1}},
		},
		{
			name: "树高为两层",
			args: args{s: []int{1, 2, 3}},
			want: [][]int{{1, 2, 3}, {2, 1, 3}, {2, 3, 1}},
		},
		{
			name: "树高为三层，且含右NULL节点",
			args: args{s: []int{1, 2, 3, NULL, 4, 5}},
			want: [][]int{{1, 2, 4, 3, 5}, {2, 4, 1, 5, 3}, {4, 2, 5, 3, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myTest(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("myTest() = %v, want %v", got, tt.want)
			}
		})
	}
}
