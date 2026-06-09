// https://leetcode.com/problems/maximum-total-subarray-value-i/
package maximumtotalsubarrayvaluei

import "math"

func maxTotalValue(nums []int, k int) int64 {
	m1 := math.MaxInt
	m2 := math.MinInt
	for i := range nums {
		m1 = min(m1, nums[i])
		m2 = max(m2, nums[i])
	}
	return int64(k) * int64(m2-m1)
}

var MaxTotalValue = maxTotalValue
