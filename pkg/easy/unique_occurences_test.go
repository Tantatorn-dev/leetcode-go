package easy

import "testing"

func Test_uniqueOccurrences(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "unique occurences",
			args: args{
				arr: []int{1, 2, 2, 1, 1, 3},
			},
			want: true,
		},
		{
			name: "no unique occurences",
			args: args{
				arr: []int{1, 1, 2, 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueOccurrences(tt.args.arr); got != tt.want {
				t.Errorf("uniqueOccurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}
