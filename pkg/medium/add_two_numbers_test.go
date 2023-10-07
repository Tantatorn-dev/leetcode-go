package medium

import (
	"reflect"
	"testing"

	"leetcode-go/pkg/common"
)

func TestAddTwoNumbers(t *testing.T) {
	type args struct {
		l1 *common.ListNode
		l2 *common.ListNode
	}
	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{
			name: "ez case",
			args: args{
				l1: common.ConstructList([]int{1, 2}),
				l2: common.ConstructList([]int{3, 5}),
			},
			want: common.ConstructList([]int{4, 7}),
		},
		{
			name: "len l1 > len l2",
			args: args{
				l1: common.ConstructList([]int{2, 4, 3}),
				l2: common.ConstructList([]int{5, 6}),
			},
			want: common.ConstructList([]int{7, 0, 4}),
		},
		{
			name: "len l1 < len l2",
			args: args{
				l1: common.ConstructList([]int{2, 4}),
				l2: common.ConstructList([]int{9, 6, 1}),
			},
			want: common.ConstructList([]int{1, 1, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
