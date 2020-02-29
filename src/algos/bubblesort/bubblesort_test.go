package algo

import (
	"reflect"
	"testing"
)

func Test_bubbleSort(t *testing.T) {
	tests := []struct {
		list []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{4, 6, 5, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 4, 3, 6, 5}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{6, 2, 4, 3, 1, 5}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		if got := bubbleSort(tt.list); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("bubbleSort() = %v, want %v", got, tt.want)
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	in := []int{6, 5, 4, 3, 2, 1}
	for i := 0; i < b.N; i++ {
		bubbleSort(in)
	}
}
