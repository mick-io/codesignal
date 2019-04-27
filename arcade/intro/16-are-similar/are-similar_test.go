package codesignal

import "testing"

func Test_areSimilar(t *testing.T) {
	tests := []struct {
		a, b   []int
		expect bool
	}{
		{
			a:      []int{1, 2, 3},
			b:      []int{1, 2, 3},
			expect: true,
		},
		{
			a:      []int{1, 2, 3},
			b:      []int{2, 1, 3},
			expect: true,
		},
		{
			a:      []int{1, 2, 2},
			b:      []int{2, 1, 1},
			expect: false,
		},
		{
			a:      []int{1, 1, 4},
			b:      []int{1, 2, 3},
			expect: false,
		},
		{
			a:      []int{1, 2, 3},
			b:      []int{1, 10, 2},
			expect: false,
		},
		{
			a:      []int{2, 3, 1},
			b:      []int{1, 3, 2},
			expect: true,
		},
		{
			a:      []int{2, 3, 9},
			b:      []int{10, 3, 2},
			expect: false,
		},
		{
			a:      []int{4, 6, 3},
			b:      []int{3, 4, 6},
			expect: false,
		},
		{
			a:      []int{832, 998, 148, 570, 533, 561, 894, 147, 455, 279},
			b:      []int{832, 998, 148, 570, 533, 561, 455, 147, 894, 279},
			expect: true,
		},
		{
			a:      []int{832, 998, 148, 570, 533, 561, 894, 147, 455, 279},
			b:      []int{832, 570, 148, 998, 533, 561, 455, 147, 894, 279},
			expect: false,
		},
	}
	for _, test := range tests {
		out := areSimilar(test.a, test.b)
		if out != test.expect {
			msg := "\n[FAIL] areSimilar(%v, %v) == %t expect %t"
			t.Errorf(msg, test.a, test.b, out, test.expect)
		}
	}
}
