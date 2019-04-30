package codesignal

import "testing"

func Test_avoidObstacle(t *testing.T) {
	tests := []struct {
		in     []int
		expect int
	}{
		{
			in:     []int{5, 3, 6, 7, 9},
			expect: 4,
		},
		{
			in:     []int{2, 3},
			expect: 4,
		},
		{
			in:     []int{1, 4, 10, 6, 2},
			expect: 7,
		},
		{
			in:     []int{1000, 999},
			expect: 6,
		},
		{
			in:     []int{19, 32, 11, 23},
			expect: 3,
		},
		{
			in:     []int{5, 8, 9, 13, 14},
			expect: 6,
		},
	}
	for _, test := range tests {
		out := avoidObstacles(test.in)
		if out != test.expect {
			msg := "\n[FAIL] avoidObstacles(%v) == %d expect %d"
			t.Errorf(msg, test.in, out, test.expect)
		}
	}
}
