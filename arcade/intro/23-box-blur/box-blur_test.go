package codesignal

import "testing"

func Test_boxBlur(t *testing.T) {
	tests := []struct{ in, expect [][]int }{
		{
			in: [][]int{
				[]int{1, 1, 1},
				[]int{1, 7, 1},
				[]int{1, 1, 1},
			},
			expect: [][]int{
				[]int{1},
			},
		},
		{
			in: [][]int{
				[]int{0, 18, 9},
				[]int{27, 9, 0},
				[]int{81, 63, 45},
			},
			expect: [][]int{
				[]int{28},
			},
		},
		{
			in: [][]int{
				[]int{36, 0, 18, 9},
				[]int{27, 54, 9, 0},
				[]int{81, 63, 72, 45},
			},
			expect: [][]int{
				[]int{40, 30},
			},
		},
		{
			in: [][]int{
				[]int{7, 4, 0, 1},
				[]int{5, 6, 2, 2},
				[]int{6, 10, 7, 8},
				[]int{1, 4, 2, 0},
			},
			expect: [][]int{
				[]int{5, 4},
				[]int{4, 4},
			},
		},
		{
			in: [][]int{
				[]int{36, 0, 18, 9, 9, 45, 27},
				[]int{27, 0, 54, 9, 0, 63, 90},
				[]int{81, 63, 72, 45, 18, 27, 0},
				[]int{0, 0, 9, 81, 27, 18, 45},
				[]int{45, 45, 27, 27, 90, 81, 72},
				[]int{45, 18, 9, 0, 9, 18, 45},
				[]int{27, 81, 36, 63, 63, 72, 81},
			},
			expect: [][]int{
				[]int{39, 30, 26, 25, 31},
				[]int{34, 37, 35, 32, 32},
				[]int{38, 41, 44, 46, 42},
				[]int{22, 24, 31, 39, 45},
				[]int{37, 34, 36, 47, 59},
			},
		},
		{
			in: [][]int{
				[]int{36, 0, 18, 9, 9, 45, 27},
				[]int{27, 0, 254, 9, 0, 63, 90},
				[]int{81, 255, 72, 45, 18, 27, 0},
				[]int{0, 0, 9, 81, 27, 18, 45},
				[]int{45, 45, 227, 227, 90, 81, 72},
				[]int{45, 18, 9, 255, 9, 18, 45},
				[]int{27, 81, 36, 127, 255, 72, 81},
			},
			expect: [][]int{
				[]int{82, 73, 48, 25, 31},
				[]int{77, 80, 57, 32, 32},
				[]int{81, 106, 88, 68, 42},
				[]int{44, 96, 103, 89, 45},
				[]int{59, 113, 137, 126, 80},
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
		output := boxBlur(test.in)
		if !expectedOutput(output, test.expect) {
			msg := "\n[FAIL] boxBlur(%v) == %v expect %v"
			t.Errorf(msg, test.in, output, test.expect)
		}
	}
}
