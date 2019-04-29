package codesignal

import "testing"

func Test_areEquallyStrong(t *testing.T) {
	tests := []struct {
		yl, yr, fl, fr int
		expect         bool
	}{
		{yl: 10, yr: 15, fl: 15, fr: 10, expect: true},
		{yl: 15, yr: 10, fl: 15, fr: 10, expect: true},
		{yl: 15, yr: 10, fl: 15, fr: 9, expect: false},
		{yl: 10, yr: 5, fl: 5, fr: 10, expect: true},
		{yl: 10, yr: 15, fl: 5, fr: 20, expect: false},
		{yl: 10, yr: 20, fl: 10, fr: 20, expect: true},
		{yl: 5, yr: 20, fl: 20, fr: 5, expect: true},
		{yl: 20, yr: 15, fl: 5, fr: 20, expect: false},
		{yl: 5, yr: 10, fl: 5, fr: 10, expect: true},
		{yl: 1, yr: 10, fl: 10, fr: 0, expect: false},
		{yl: 5, yr: 5, fl: 10, fr: 10, expect: false},
		{yl: 10, yr: 5, fl: 10, fr: 0, expect: false},
		{yl: 1, yr: 1, fl: 1, fr: 1, expect: true},
		{yl: 0, yr: 10, fl: 10, fr: 0, expect: true},
	}
	for _, test := range tests {
		out := areEquallyStrong(test.yl, test.yr, test.fl, test.fr)
		if out != test.expect {
			msg := "\n[FAIL] areEquallyStrong(%d, %d, %d, %d) == %t expect %t"
			t.Errorf(msg, test.yl, test.yr, test.fl, test.fr, out, test.expect)
		}
	}
}
