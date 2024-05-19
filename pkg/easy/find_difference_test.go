package easy

import (
	"reflect"
	"testing"
)

func Test_findDifference(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "No intersected values",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{4, 5},
			},
			want: [][]int{{1, 2, 3}, {4, 5}},
		},
		{
			name: "Has intersected values",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{2, 5},
			},
			want: [][]int{{1, 3}, {5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDifference(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
