package medium

import (
	"leetcode-go/pkg/common"
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *common.ListNode
	}
	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{
			name: "[1,2,3,3,4,4,5]",
			args: args{
				head: common.ConstructList([]int{1, 2, 3, 3, 4, 4, 5}),
			},
			want: common.ConstructList([]int{1, 2, 5}),
		},
		{
			name: "[1,1,1,2,3]",
			args: args{
				head: common.ConstructList([]int{1, 1, 1, 2, 3}),
			},
			want: common.ConstructList([]int{2, 3}),
		},
		{
			name: "[1,1]",
			args: args{
				head: common.ConstructList([]int{1, 1, 1, 2, 3}),
			},
			want: common.ConstructList([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
