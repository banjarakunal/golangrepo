package algos

import "testing"

func Test_highestProduct(t *testing.T) {

	tests := []struct {
		in []int
		want int
	}{
		{[]int{-10,-10,1,3,2}, 300},
		{[]int{20,20,1,1,1}, 400},
	}
	for _, tt := range tests {
			if got := highestProduct(tt.in); got != tt.want {
				t.Errorf("highestProduct() = %v, want %v", got, tt.want)
			}
	}
}
