package common

import (
	"testing"
)

func Test_Max(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{[]int{1, 2, 3}, 3},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{1, 2}, 2},
	}
	for _, tt := range tests {
		if got := Max(tt.in...); got != tt.want {
			t.Errorf("max() = %v, want %v", got, tt.want)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{[]int{1, 2, 3}, 1},
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{1, 2}, 1},
	}
	for _, tt := range tests {
		if got := Min(tt.in...); got != tt.want {
			t.Errorf("max() = %v, want %v", got, tt.want)
		}
	}
}
