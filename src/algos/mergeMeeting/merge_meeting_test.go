package main

import (
	"reflect"
	"testing"
)

func Test_mergedMeetings(t *testing.T) {
	tests := []struct {
		in []meeting
		want []meeting
	}{
		{[]meeting{}, []meeting{}},
		{[]meeting{{1, 2}}, []meeting{{1, 2}}},
		{[]meeting{{1, 2}, {2, 3}}, []meeting{{1, 3}}},
		{[]meeting{{1, 5}, {2, 3}}, []meeting{{1, 5}}},
		{[]meeting{{1, 2}, {4, 5}}, []meeting{{1, 2}, {4, 5}}},
		{[]meeting{{1, 5}, {2, 3}, {4, 5}}, []meeting{{1, 5}}},
		{[]meeting{{1, 2}, {2, 3}, {4, 5}}, []meeting{{1, 3}, {4, 5}}},
		{[]meeting{{1, 6}, {2, 3}, {4, 5}}, []meeting{{1, 6}}},
		{[]meeting{{4, 5}, {2, 3}, {1, 6}}, []meeting{{1, 6}}},
	}
	for _, tt := range tests {
			if got := mergedMeetings(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergedMeetings() = %v, want %v", got, tt.want)
			}
		
	}
}
