package codesignal

import (
	"testing"
)

func Test_almostIncreaseSequence(t *testing.T) {
	tests := []struct {
		input  []int
		expect bool
	}{
		{[]int{1, 3, 2, 1}, false},
		{[]int{1, 3, 2}, true},
		{[]int{1, 2, 1, 2}, false},
		{[]int{3, 6, 5, 8, 10, 20, 15}, false},
		{[]int{1, 1, 2, 3, 4, 4}, false},
		{[]int{1, 4, 10, 4, 2}, false},
		{[]int{10, 1, 2, 3, 4, 5}, true},
		{[]int{1, 1, 1, 2, 3}, false},
		{[]int{0, -2, 5, 6}, true},
		{[]int{1, 2, 3, 4, 5, 3, 5, 6}, false},
		{[]int{40, 50, 60, 10, 20, 30}, false},
		{[]int{1, 1}, true},
		{[]int{1, 2, 5, 3, 5}, true},
		{[]int{1, 2, 5, 5, 5}, false},
		{[]int{10, 1, 2, 3, 4, 5, 6, 1}, false},
		{[]int{1, 2, 3, 4, 3, 6}, true},
		{[]int{1, 2, 3, 4, 99, 5, 6}, true},
		{[]int{123, -17, -5, 1, 2, 3, 12, 43, 45}, true},
		{[]int{3, 5, 67, 98, 3}, true},
	}
	for _, test := range tests {
		original := make([]int, len(test.input))
		copy(original, test.input)
		output := almostIncreasingSequence(test.input)
		if output != test.expect {
			msg := "\n[FAIL] almostIncreasingSequence(%v) == %t expect %t"
			t.Errorf(msg, original, output, test.expect)
		}
	}
}
