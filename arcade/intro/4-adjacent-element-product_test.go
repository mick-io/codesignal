package codesignal

import (
	"testing"
)

func Test_adjacentElementProduct(t *testing.T) {
	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{1, 2, 3, 4, 5}, 20},
		{[]int{3, 6, -2, -5, 7, 3}, 21},
		{[]int{-1, -2}, 2},
		{[]int{5, 1, 2, 3, 1, 4}, 6},
	}
	for _, test := range tests {
		output := adjacentElementsProduct(test.input)
		if output != test.expect {
			msg := "\n[FAIL] adjacentElementsProduct(%v) == %d expect %d"
			t.Errorf(msg, test.input, output, test.expect)
		}
	}
}
