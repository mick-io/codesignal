package codesignal

import "testing"

func Test_sortByHeight(t *testing.T) {
	tests := []struct{ in, expect []int }{
		{
			in:     []int{-1, 150, 190, 170, -1, -1, 160, 180},
			expect: []int{-1, 150, 160, 170, -1, -1, 180, 190},
		},
		{
			in:     []int{-1, -1, -1, -1, -1},
			expect: []int{-1, -1, -1, -1, -1},
		},
		{
			in:     []int{-1},
			expect: []int{-1},
		},
		{
			in:     []int{4, 2, 9, 11, 2, 16},
			expect: []int{2, 2, 4, 9, 11, 16},
		},
		{
			in:     []int{2, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 1},
			expect: []int{1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 2},
		},
		{
			in:     []int{23, 54, -1, 43, 1, -1, -1, 77, -1, -1, -1, 3},
			expect: []int{1, 3, -1, 23, 43, -1, -1, 54, -1, -1, -1, 77},
		},
	}
	for _, test := range tests {
		out := sortByHeight(test.in)
		if !slicesMatch(out, test.expect) {
			msg := "\n[FAIL] sortByHeight(%v) == %v expect %v"
			t.Errorf(msg, test.in, out, test.expect)
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
