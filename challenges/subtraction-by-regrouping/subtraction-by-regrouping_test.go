package codesignal

import "testing"

func Test_subtractionByRegrouping(t *testing.T) {
	tests := []struct {
		minuend, subtrahend int
		expect              []int
	}{
		{
			minuend:    2024,
			subtrahend: 1234,
			expect:     []int{4, 12, 9, 1},
		},
		{
			minuend:    6,
			subtrahend: 5,
			expect:     []int{6},
		},
		{
			minuend:    4567,
			subtrahend: 3456,
			expect:     []int{7, 6, 5, 4},
		},
		{
			minuend:    3001,
			subtrahend: 2901,
			expect:     []int{1, 0, 10, 2},
		},
		{
			minuend:    5000,
			subtrahend: 4999,
			expect:     []int{10, 9, 9, 4},
		},
		{
			minuend:    1,
			subtrahend: 1,
			expect:     []int{1},
		},
		{
			minuend:    51234,
			subtrahend: 12345,
			expect:     []int{14, 12, 11, 10, 4},
		},
		{
			minuend:    20100,
			subtrahend: 19199,
			expect:     []int{10, 9, 10, 9, 1},
		},
	}
	slicesMatch := func(s1, s2 []int) bool {
		for i, n := range s1 {
			if n != s2[i] {
				return false
			}
		}
		return true
	}
	for _, test := range tests {
		out := subtractionByRegrouping(test.minuend, test.subtrahend)
		if !slicesMatch(out, test.expect) {
			m := "[FAIL] subtractionByRegrouping(%d, %d) == %v expect %v"
			t.Errorf(m, test.minuend, test.subtrahend, out, test.expect)
		}
	}
}
