package codesignal

import "testing"

func Test_alternatingSums(t *testing.T) {
	test := []struct{ in, expect []int }{
		{
			in:     []int{50, 60, 60, 45, 70},
			expect: []int{180, 105},
		},
		{
			in:     []int{100, 50},
			expect: []int{100, 50},
		},
		{
			in:     []int{80},
			expect: []int{80, 0},
		},
		{
			in:     []int{100, 50, 50, 100},
			expect: []int{150, 150},
		},
		{
			in:     []int{100, 51, 50, 100},
			expect: []int{150, 151},
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
	for _, test := range test {
		out := alternatingSums(test.in)
		if !slicesMatch(test.expect, out) {
			msg := "\n[FAIL] alternatingSums(%v) == %v expect %v"
			t.Errorf(msg, test.in, out, test.expect)
		}
	}
}
