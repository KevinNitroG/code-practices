// https://leetcode.com/problems/partition-array-according-to-given-pivot/

package main

import "slices"

func pivotArray(nums []int, pivot int) []int {
	less := []int{}
	greater := []int{}
	pivotAppearances := 0

	for _, val := range nums {
		switch {
		case val < pivot:
			less = append(less, val)
		case val > pivot:
			greater = append(greater, val)
		default:
			pivotAppearances++
		}
	}
	pivots := make([]int, pivotAppearances)
	for i := range pivotAppearances {
		pivots[i] = pivot
	}
	return slices.Concat(less, pivots, greater)
}
