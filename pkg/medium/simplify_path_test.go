package medium

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "/home/",
			args: args{
				path: "/home/",
			},
			want: "/home",
		},
		{
			name: "/home//foo",
			args: args{
				path: "/home//foo",
			},
			want: "/home/foo",
		},
		{
			name: "/../",
			args: args{
				path: "/../",
			},
			want: "/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
