package main

import (
	"reflect"
	"testing"
)

func Test_reverse(t *testing.T) {
	tests := []struct {
		in []string
		want []string
	}{
		{[]string{}, []string{}},
		{[]string{"a"}, []string{"a"}},
		{[]string{"a", "b"}, []string{"b", "a"}},
		{[]string{"a", "b", "c"}, []string{"c", "b", "a"}},
		{[]string{"a", "b", "c", "d"}, []string{"d", "c", "b", "a"}},
	}
	for _, tt := range tests {
			if got := reverse(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
	}
}
