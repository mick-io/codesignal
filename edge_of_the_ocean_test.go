package codesignal

import "testing"

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
			msg := "\n[FAIL] adjacentElementsProduct(%s) == %d expect %d"
			t.Errorf(msg, test.input, output, test.expect)
		}
	}
}

func Test_shapeArea(t *testing.T) {
	tests := []struct{ input, expect int }{
		{2, 5},
		{3, 13},
		{1, 1},
		{5, 41},
		{7000, 97986001},
	}
	for _, test := range tests {
		output := shapeArea(test.input)
		if output != test.expect {
			msg := "\n[FAIL] shapeArea(%d) == %d expect %d"
			t.Errorf(msg, test.input, output, test.expect)
		}
	}
}
