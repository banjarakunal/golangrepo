package main

import "sort"

func mergSortedSlice(s1 []int, s2 []int)[]int{
	out := []int{}
	out = append(out, s1...)
	out = append(out, s2...)
	sort.Slice(out, func(i, j int)bool{
		return out[i] < out[j]
	})
	return out
}