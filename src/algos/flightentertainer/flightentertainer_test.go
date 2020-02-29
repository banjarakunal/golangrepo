package main

import "testing"

func Test_flightentertrainer(t *testing.T) {

	tests := []struct {
		movLen []int
		total  int
		want   bool
	}{
		{[]int{}, 1, false},
		{[]int{0}, 1, false},
		{[]int{1}, 1, false},
		{[]int{0, 1}, 1, true},
		{[]int{1, 1}, 2, true},
		{[]int{2, 3, 4}, 6, true},
		{[]int{2, 3, 4}, 8, false},
	}
	for _, tt := range tests {
		if got := flightentertrainer(tt.movLen, tt.total); got != tt.want {
			t.Errorf("flightentertrainer() = %v, want %v", got, tt.want)
		}
	}
}
