package codesignal

import "testing"

func Test_improperFractionToMixed(t *testing.T) {
	tests := []struct{ input, expect []int }{
		{
			input:  []int{7, 2},
			expect: []int{3, 1, 2},
		},
		{
			input:  []int{10, 3},
			expect: []int{3, 1, 3},
		},
		{
			input:  []int{23, 22},
			expect: []int{1, 1, 22},
		},
		{
			input:  []int{7, 3},
			expect: []int{2, 1, 3},
		},
		{
			input:  []int{8, 5},
			expect: []int{1, 3, 5},
		},
		{
			input:  []int{15, 4},
			expect: []int{3, 3, 4},
		},
		{
			input:  []int{18, 7},
			expect: []int{2, 4, 7},
		},
	}
	for _, test := range tests {
		output := improperFractionToMixed(test.input)
		if !slicesMatch(test.expect, output) {
			msg := "\n[FAIL] improperFractionToMixed(%v) == %v expect %v"
			t.Errorf(msg, test.input, output, test.expect)
		}
	}
}

func slicesMatch(s1, s2 []int) bool {
	for i, n := range s1 {
		if n != s2[i] {
			return false
		}
	}
	return true
}
