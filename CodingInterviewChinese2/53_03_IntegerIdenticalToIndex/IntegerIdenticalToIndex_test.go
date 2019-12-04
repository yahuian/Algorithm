package _53_03

import "testing"

func TestIntegerIdenticalToIndex(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				arr: []int{-3, -1, 1, 3, 5},
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				arr: []int{0, 1, 3, 5, 6},
			},
			want: 0,
		},
		{
			name: "test3",
			args: args{
				arr: []int{-1, 0, 1, 2, 4},
			},
			want: 4,
		},
		{
			name: "test4",
			args: args{
				arr: []int{-1, 0, 1, 2, 5},
			},
			want: -1,
		},
		{
			name: "test5",
			args: args{
				arr: []int{0},
			},
			want: 0,
		},
		{
			name: "test6",
			args: args{
				arr: []int{10},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntegerIdenticalToIndex(tt.args.arr); got != tt.want {
				t.Errorf("IntegerIdenticalToIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
