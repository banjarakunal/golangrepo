package main

import (
	"testing"
)

func TestStack_Size(t *testing.T) {
	tests := []struct {
		s    *Stack
		want int
	}{
		{&Stack{list: []int{}}, 0},
		{&Stack{list: []int{1}}, 1},
		{&Stack{list: []int{1, 2, 3, 4}}, 4},
	}
	for _, tt := range tests {
		if got := tt.s.Size(); got != tt.want {
			t.Errorf("Stack.Size() = %v, want %v", got, tt.want)
		}
	}
}

func TestStack_Push(t *testing.T) {
	st := Stack{
		list: []int{},
	}

	st.Push(1)
	st.Push(2)
	st.Push(3)
	if st.Size() != 3 {
		t.Errorf("Stack.Push() = %v, want %v", st.Size(), 3)
	}
}

func TestStack_Pop(t *testing.T) {
	st := Stack{
		list: []int{},
	}

	st.Push(1)
	st.Push(2)
	st.Push(3)

	st.Pop()
	if st.Size() != 2 {
		t.Errorf("Stack.Push() = %v, want %v", st.Size(), 3)
	}
}
func TestStack_IsEmpty(t *testing.T) {
	tests := []struct {
		s    *Stack
		want bool
	}{
		{&Stack{list: []int{}}, true},
		{&Stack{list: []int{1}}, false},
		{&Stack{list: []int{1, 2, 3, 4}}, false},
	}
	for _, tt := range tests {
		if got := tt.s.IsEmpty(); got != tt.want {
			t.Errorf("Stack.Size() = %v, want %v", got, tt.want)
		}
	}
}

func BenchmarkPush(b *testing.B) {
	st := Stack{
		list: []int{},
	}
	for n := 0; n < b.N; n++ {
		st.Push(10)
	}
}
