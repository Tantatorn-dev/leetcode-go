package easy

import "testing"

func TestStrStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "found needle",
			args: args{
				haystack: "leetcode",
				needle:   "code",
			},
			want: 4,
		},
		{
			name: "not found needle",
			args: args{
				haystack: "leetcode",
				needle:   "leeto",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("StrStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
