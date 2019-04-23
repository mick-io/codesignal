package codesignal

import (
	"testing"
)

func Test_matrixElementSum(t *testing.T) {
	tests := []struct {
		in     [][]int
		expect int
	}{
		{
			in: [][]int{
				[]int{0, 1, 1, 2},
				[]int{0, 5, 0, 0},
				[]int{2, 0, 3, 3},
			},
			expect: 9,
		},
		{
			in: [][]int{
				[]int{1, 1, 1, 0},
				[]int{0, 5, 0, 1},
				[]int{2, 1, 3, 10},
			},
			expect: 9,
		},
		{
			in: [][]int{
				[]int{1, 1, 1},
				[]int{2, 2, 2},
				[]int{3, 3, 3},
			},
			expect: 18,
		},
		{
			in: [][]int{
				[]int{0},
			},
			expect: 0,
		},
		{
			in: [][]int{
				[]int{1, 0, 3},
				[]int{0, 2, 1},
				[]int{1, 2, 0},
			},
			expect: 5,
		},
		{
			in: [][]int{
				[]int{1},
				[]int{5},
				[]int{0},
				[]int{2},
			},
			expect: 6,
		},
		{
			in: [][]int{
				[]int{1, 2, 3, 4, 5},
			},
			expect: 15,
		},
		{
			in: [][]int{
				[]int{2},
				[]int{5},
				[]int{10},
			},
			expect: 17,
		},
		{
			in: [][]int{
				[]int{4, 0, 1},
				[]int{10, 7, 0},
				[]int{0, 0, 0},
				[]int{9, 1, 2},
			},
			expect: 15,
		},
		{
			in: [][]int{
				[]int{1},
			},
			expect: 1,
		},
	}
	for _, test := range tests {
		out := matrixElementsSum(test.in)
		if out != test.expect {
			msg := "\n[FAIL] matrixElementSum(%v) == % expect %"
			t.Errorf(msg, test.in, out, test.expect)
		}
	}
}
