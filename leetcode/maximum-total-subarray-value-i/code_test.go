package maximumtotalsubarrayvaluei_test

import (
	"testing"

	maximumtotalsubarrayvaluei "github.com/KevinNitroG/code-practices/leetcode/maximum-total-subarray-value-i"
	"github.com/stretchr/testify/assert"
)

func TestMaxTotalValue(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		nums []int
		k    int
		res  int64
	}{
		{
			name: "First",
			nums: []int{1, 3, 2},
			k:    2,
			res:  4,
		},
		{
			name: "Second",
			nums: []int{4, 2, 5, 1},
			k:    3,
			res:  12,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := maximumtotalsubarrayvaluei.MaxTotalValue(tc.nums, tc.k)
			assert.Equal(t, tc.res, res)
		})
	}
}
