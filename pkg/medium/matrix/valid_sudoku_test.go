package matrix

import "testing"

func Test_isRowValid(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid",
			args: args{
				board: [][]byte{
					{'1', '.', '2', '.', '.', '.', '.'},
				},
			},
			want: true,
		},
		{
			name: "invalid",
			args: args{
				board: [][]byte{
					{'1', '.', '2', '.', '2', '.', '.'},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRowValid(tt.args.board); got != tt.want {
				t.Errorf("isRowValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSubBoxesValid(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid",
			args: args{
				board: [][]byte{
					{'1', '2', '3', '4', '5', '6', '7', '8', '9'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				},
			},
			want: true,
		},
		{
			name: "invalid",
			args: args{
				board: [][]byte{
					{'1', '2', '3', '4', '5', '6', '7', '8', '9'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '2', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubBoxesValid(tt.args.board); got != tt.want {
				t.Errorf("isSubBoxesValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidSudoku(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid",
			args: args{
				board: [][]byte{
					{'1', '2', '3', '4', '5', '6', '7', '8', '9'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				},
			},
			want: true,
		},
		{
			name: "invalid",
			args: args{
				board: [][]byte{
					{'1', '2', '3', '4', '5', '6', '7', '8', '9'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '2', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidSudoku(tt.args.board); got != tt.want {
				t.Errorf("isValidSudoku() = %v, want %v", got, tt.want)
			}
		})
	}
}
