package linkedlist

import (
	"leetcode-go/pkg/common"
	"reflect"
	"testing"
)

func Test_deleteMiddle(t *testing.T) {
	type args struct {
		head *common.ListNode
	}
	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{
			name: "length = 7",
			args: args{
				head: common.ConstructList([]int{1, 3, 4, 7, 1, 2, 6}),
			},
			want: common.ConstructList([]int{1, 3, 4, 1, 2, 6}),
		},
		{
			name: "length = 4",
			args: args{
				head: common.ConstructList([]int{1, 2, 3, 4}),
			},
			want: common.ConstructList([]int{1, 2, 4}),
		},
		{
			name: "length = 0",
			args: args{
				head: common.ConstructList([]int{1}),
			},
			want: common.ConstructList([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteMiddle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteMiddle() = %v, want %v", got, tt.want)
			}
		})
	}
}
