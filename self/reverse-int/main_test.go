package reverseint_test

import (
	"testing"

	reverseint "github.com/KevinNitroG/code-practices/self/reverse-int"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		in   int
		res  int
	}{
		{
			name: "Positive",
			in:   1234,
			res:  4321,
		},
		{
			name: "Has zero in the middle",
			in:   10034,
			res:  43001,
		},
		{
			name: "Zero at the end",
			in:   524300,
			res:  3425,
		},
		{
			name: "Negative",
			in:   -1258,
			res:  -8521,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := reverseint.Reverse(tc.in)
			assert.Equal(t, tc.res, res)
		})
	}
}
