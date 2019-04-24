package codesignal

import (
	"testing"
)

func Test_centuryFromYear(t *testing.T) {
	tests := []struct{ input, expect int }{
		{1905, 20},
		{1700, 17},
		{200, 2},
		{45, 1},
	}
	for _, test := range tests {
		output := centuryFromYear(test.input)
		if output != test.expect {
			msg := "\n[FAIL] centuryFromYear(%d) == %d expect %d"
			t.Errorf(msg, test.input, output, test.expect)
		}
	}
}
