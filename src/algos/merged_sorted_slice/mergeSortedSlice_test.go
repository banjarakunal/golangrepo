package main

import (
	"reflect"
	"testing"
)

func Test_mergSortedSlice(t *testing.T) {
	
	tests := []struct {
		in1 []int
		in2 []int
		want []int
	}{
		{
			[]int{},
			[]int{},
			[]int{},
		},
		{
			[]int{},
			[]int{1},
			[]int{1},
		},
		{
			[]int{1, 3, 5},
			[]int{2, 4, 6},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			[]int{1, 3, 5},
			[]int{2, 4, 6, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{6},
			[]int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
			if got := mergSortedSlice(tt.in1, tt.in2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergSortedSlice() = %v, want %v", got, tt.want)
			}
	}
}
