package codesignal

import "testing"

func Test_arrayMaximalAdjacentDifference(t *testing.T) {
	tests := []struct {
		in     []int
		expect int
	}{
		{
			in:     []int{2, 4, 1, 0},
			expect: 3,
		},
		{
			in:     []int{1, 1, 1, 1},
			expect: 0,
		},
		{
			in:     []int{-1, 4, 10, 3, -2},
			expect: 7,
		},
		{
			in:     []int{10, 11, 13},
			expect: 2,
		},
		{
			in:     []int{-2, -2, -2, -2, -2},
			expect: 0,
		},
		{
			in:     []int{-1, 1, -3, -4},
			expect: 4,
		},
	}
	for _, test := range tests {
		out := arrayMaximalAdjacentDifference(test.in)
		if out != test.expect {
			msg := "\n[FAIL] arrayMaximalAdjacentDifference(%v) == %d expect %d"
			t.Errorf(msg, test.in, out, test.expect)
		}
	}
}
