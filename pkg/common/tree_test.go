package common

import (
	"reflect"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}

	simpleTree := new(TreeNode)
	simpleTree.Val = 2
	simpleTree.Left = new(TreeNode)
	simpleTree.Left.Val = 1
	simpleTree.Right = new(TreeNode)
	simpleTree.Right.Val = 3

	oneNodeTree := new(TreeNode)
	oneNodeTree.Val = 2

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "one node",
			args: args{
				nums: []int{2},
			},
			want: oneNodeTree,
		},
		{
			name: "simple",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: simpleTree,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortedArrayToBST(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortedArrayToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "max_depth",
			args: args{
				root: SortedArrayToBST([]int{1, 2, 3}),
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
