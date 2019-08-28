package _0844

import "testing"

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		S string
		T string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// 单元测试
		{
			name: "test1",
			args: args{
				S: "ab#c",
				T: "ad#c",
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				S: "ab##",
				T: "c#d#",
			},
			want: true,
		},
		{
			name: "test3",
			args: args{
				S: "a##c",
				T: "#a#c",
			},
			want: true,
		},
		{
			name: "test4",
			args: args{
				S: "a#c",
				T: "b",
			},
			want: false,
		},
		{
			name: "test5",
			args: args{
				S: "xywrrmp",
				T: "xywrrmu#p",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.args.S, tt.args.T); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
