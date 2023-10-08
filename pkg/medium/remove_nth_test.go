package medium

import (
	"reflect"
	"testing"

	"leetcode-go/pkg/common"
)

func TestRemoveNthFromEnd(t *testing.T) {
	type args struct {
		head *common.ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{
			name: "removed success",
			args: args{
				head: common.ConstructList([]int{1, 2, 3, 4}),
				n:    1,
			},
			want: common.ConstructList([]int{1, 2, 3}),
		},
		{
			name: "removed success",
			args: args{
				head: common.ConstructList([]int{1, 2, 3, 4, 5}),
				n:    3,
			},
			want: common.ConstructList([]int{1, 2, 4, 5}),
		},
		{
			name: "removed last element",
			args: args{
				head: common.ConstructList([]int{1, 2, 3, 4}),
				n:    1,
			},
			want: common.ConstructList([]int{1, 2, 3}),
		},
		{
			name: "removed from only one",
			args: args{
				head: common.ConstructList([]int{1}),
				n:    1,
			},
			want: common.ConstructList([]int{}),
		},
		{
			name: "removed from two",
			args: args{
				head: common.ConstructList([]int{1, 2}),
				n:    2,
			},
			want: common.ConstructList([]int{2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
