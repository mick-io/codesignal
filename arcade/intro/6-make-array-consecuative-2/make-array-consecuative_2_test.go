package codesignal

import (
	"testing"
)

func Test_makeArrayConsecutive2(t *testing.T) {
	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{6, 2, 3, 8}, 3},
		{[]int{0, 3}, 2},
		{[]int{5, 4, 6}, 0},
		{[]int{6, 3}, 2},
	}
	for _, test := range tests {
		output := makeArrayConsecutive2(test.input)
		if output != test.expect {
			msg := "\n[FAIL] makeArrayConsecutive2(%v) == %d expect %d"
			t.Errorf(msg, test.input, output, test.expect)
		}
	}
}
