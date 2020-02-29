package algos

import (
	"testing"
	"reflect"
)

func Test_getProductOfOthers(t *testing.T) {
	tests := []struct {
		in []int
		want []int
	}{
		{[]int{1, 7, 3, 4}, []int{84, 12, 28, 21}},
		{[]int{2, 4, 10}, []int{40, 20, 8}},
		{[]int{2, 4, 0}, []int{0, 0, 8}},
	}

	for _, tt := range tests {
		if got := getProductOfOthers(tt.in); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("highestProduct() = %v, want %v", got, tt.want)
		}
	}
}
