package common

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sorted already",
			args: args{
				arr: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "not sorted yet",
			args: args{
				arr: []int{4, 2, 1, 3, 6},
			},
			want: []int{1, 2, 3, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sorted already",
			args: args{
				arr: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "not sorted yet",
			args: args{
				arr: []int{4, 2, 1, 3, 6},
			},
			want: []int{1, 2, 3, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectionSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
