package codesignal

import "testing"

func Test_isIPv4Address(t *testing.T) {
	tests := []struct {
		in     string
		expect bool
	}{
		{
			in:     "172.16.254.1",
			expect: true,
		},
		{
			in:     "172.316.254.1",
			expect: false,
		},
		{
			in:     ".254.255.0",
			expect: false,
		},
		{
			in:     "1.1.1.1a",
			expect: false,
		},
		{
			in:     "1",
			expect: false,
		},
		{
			in:     "1.23.256.255.",
			expect: false,
		},
		{
			in:     "1.23.256..",
			expect: false,
		},
		{
			in:     "0..1.0",
			expect: false,
		},
		{
			in:     "35..36.9.9.0",
			expect: false,
		},
		{
			in:     "1.1.1.1.1",
			expect: false,
		},
		{
			in:     "1.256.1.1",
			expect: false,
		},
		{
			in:     "a0.1.1.1",
			expect: false,
		},
		{
			in:     "0.1.1.256",
			expect: false,
		},
		{
			in:     "129380129831213981.255.255.255",
			expect: false,
		},
		{
			in:     "255.255.255.255abcdekjhf",
			expect: false,
		},
		{
			in:     "7283728",
			expect: false,
		},
		{
			in:     "0..1.0.0",
			expect: false,
		},
	}
	for _, test := range tests {
		out := isIPv4Address(test.in)
		if out != test.expect {
			msg := "\n[FAIL] isIPv4Address(%s) == %t expect %t"
			t.Errorf(msg, test.in, out, test.expect)
		}
	}
}
