package algos

import (
	"reflect"
	"testing"
)

func Test_inplaceShuffle(t *testing.T) {
	tests := []struct {
		in       []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3, 4, 5}, []int{2, 1, 4, 3, 5}},
	}

	for _, tt := range tests {
			if got := inplaceShuffle(tt.in); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("inplaceShuffle() = %v, want %v", got, tt.expected)
			}
	}
}
