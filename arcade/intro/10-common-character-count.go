package codesignal

/* commonCharacterCount

Given two strings, find the number of common characters between them.

Example

For s1 = "aabcc" and s2 = "adcaa", the output should be
commonCharacterCount(s1, s2) = 3.

Strings have 3 common characters - 2 "a"s and 1 "c".

Input/Output

    [execution time limit] 4 seconds (go)

    [input] string s1

    A string consisting of lowercase English letters.

    Guaranteed constraints:
    1 ≤ s1.length < 15.

    [input] string s2

    A string consisting of lowercase English letters.

    Guaranteed constraints:
    1 ≤ s2.length < 15.

    [output] integer
*/

func commonCharacterCount(s1 string, s2 string) (nCommon int) {
	runes1, runes2 := []rune(s1), []rune(s2)
	for _, r1 := range runes1 {
		for i, r2 := range runes2 {
			if r1 == r2 {
				runes2[i] = 0
				nCommon++
				break
			}
		}
	}
	return nCommon
}
