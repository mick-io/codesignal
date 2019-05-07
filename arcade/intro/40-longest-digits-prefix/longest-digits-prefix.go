package codesignal

/* longestDigitPrefix
Given a string, output its longest prefix which contains only digits.

Example

For s = "123aa1", the output should be
longestDigitsPrefix(s) = "123".

Input/Output

    [execution time limit] 4 seconds (go)

    [input] string s

    Guaranteed constraints:
    3 ≤ s.length ≤ 100.

    [output] string
*/

// import "fmt"

func longestDigitsPrefix(s string) string {
	prefix := []rune{}
	for _, r := range s {
		if r < '0' || '9' < r {
			break
		}
		prefix = append(prefix, r)
	}
	return string(prefix)
}
