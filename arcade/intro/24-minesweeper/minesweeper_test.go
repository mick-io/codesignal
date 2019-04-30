package codesignal

import "testing"

func Test_minesweeper(t *testing.T) {
	tests := []struct {
		matrix [][]bool
		expect [][]int
	}{
		{
			matrix: [][]bool{
				[]bool{true, false, false},
				[]bool{false, true, false},
				[]bool{false, false, false},
			},
			expect: [][]int{
				[]int{1, 2, 1},
				[]int{2, 1, 1},
				[]int{1, 1, 1},
			},
		},
		{
			matrix: [][]bool{
				[]bool{false, false, false},
				[]bool{false, false, false},
			},
			expect: [][]int{
				[]int{0, 0, 0},
				[]int{0, 0, 0},
			},
		},
		{
			matrix: [][]bool{
				[]bool{true, false, false, true},
				[]bool{false, false, true, false},
				[]bool{true, true, false, true},
			},
			expect: [][]int{
				[]int{0, 2, 2, 1},
				[]int{3, 4, 3, 3},
				[]int{1, 2, 3, 1},
			},
		},
		{
			matrix: [][]bool{
				[]bool{true, true, true},
				[]bool{true, true, true},
				[]bool{true, true, true},
			},
			expect: [][]int{
				[]int{3, 5, 3},
				[]int{5, 8, 5},
				[]int{3, 5, 3},
			},
		},
		{
			matrix: [][]bool{
				[]bool{false, true, true, false},
				[]bool{true, true, false, true},
				[]bool{false, false, true, false},
			},
			expect: [][]int{
				[]int{3, 3, 3, 2},
				[]int{2, 4, 5, 2},
				[]int{2, 3, 2, 2},
			},
		},
		{
			matrix: [][]bool{
				[]bool{true, false},
				[]bool{true, false},
				[]bool{false, true},
				[]bool{false, false},
				[]bool{false, false},
			},
			expect: [][]int{
				[]int{1, 2},
				[]int{2, 3},
				[]int{2, 1},
				[]int{1, 1},
				[]int{0, 0},
			},
		},
	}
	slicesMatch := func(s1, s2 []int) bool {
		for i, n := range s1 {
			if s2[i] != n {
				return false
			}
		}
		return true
	}
	expectedOutput := func(output, expect [][]int) bool {
		for i, slice := range expect {
			if !slicesMatch(slice, output[i]) {
				return false
			}
		}
		return true
	}
	for _, test := range tests {
		out := minesweeper(test.matrix)
		if !expectedOutput(out, test.expect) {
			msg := "\n[FAIL] minesweeper(%v) == %v expect %v"
			t.Errorf(msg, test.matrix, out, test.expect)
		}
	}

}
