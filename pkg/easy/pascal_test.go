package easy

import (
	"reflect"
	"testing"
)

func Test_generatePascal(t *testing.T) {
	type args struct {
		numRows int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "5 rows",
			args: args{
				numRows: 5,
			},
			want: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			name: "3 rows",
			args: args{
				numRows: 3,
			},
			want: [][]int{{1}, {1, 1}, {1, 2, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generatePascal(tt.args.numRows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generatePascal() = %v, want %v", got, tt.want)
			}
		})
	}
}
