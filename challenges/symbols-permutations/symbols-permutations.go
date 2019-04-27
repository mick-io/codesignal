package codesignal

/* symbolsPermutation
Determine if it is possible to rearrange the characters in a string to obtain another string.

Example

    For word1 = "abc" and word2 = "cab", the output should be
    symbolsPermutation(word1, word2) = true;
    For word1 = "aaaa" and word2 = "aaa", the output should be
    symbolsPermutation(word1, word2) = false.

Input/Output

    [execution time limit] 4 seconds (go)

    [input] string word1

    String of lowercase English letters.

    Guaranteed constraints:
    3 ≤ word1.length ≤ 10.

    [input] string word2

    String of lowercase English letters.

    Guaranteed constraints:
    3 ≤ word2.length ≤ 10.

    [output] boolean
        true if it is possible to rearrange the characters in word1 to obtain word2, false otherwise.
*/

func symbolsPermutation(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	return contentsMatch([]rune(word1), []rune(word2))
}

func contentsMatch(s1, s2 []rune) bool {
	for _, r := range s1 {
		i := find(r, s2)
		if i == -1 {
			return false
		}
		s2[i] = 0
	}
	return true
}

func find(r rune, runes []rune) (i int) {
	for i, char := range runes {
		if char == r {
			return i
		}
	}
	return -1
}
