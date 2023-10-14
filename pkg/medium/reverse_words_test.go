package medium

import (
	"reflect"
	"testing"
)

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "the sky is blue",
			args: args{
				s: "the sky is blue",
			},
			want: "blue is sky the",
		},
		{
			name: "    hello   world   ",
			args: args{
				s: "     hello   world ",
			},
			want: "world hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_split(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "hello moto",
			args: args{
				s: "hello moto",
			},
			want: []string{"hello", "moto"},
		},
		{
			name: "hello moto with trailing spaces",
			args: args{
				s: "    hello moto   ",
			},
			want: []string{"hello", "moto"},
		},
		{
			name: "the sky is blue",
			args: args{
				s: "the sky is blue",
			},
			want: []string{"the", "sky", "is", "blue"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := split(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("split() = %v, want %v", got, tt.want)
			}
		})
	}
}
