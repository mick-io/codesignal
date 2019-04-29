package codesignal

/*
Given a string, delete whitespaces from it.

Example

For inputStr = "a b cd e", the output should be
deleteWhitespaces(inputStr) = "abcde".

Input/Output

    [execution time limit] 4 seconds (go)

    [input] string inputStr

    String consisting of lowercase English letters and whitespace characters.

    Guaranteed constraints:
    5 ≤ inputStr.length ≤ 15.

    [output] string
        inputString without whitespace characters.
*/
const space = ' '

func deleteWhitespaces(s string) string {
	in, out := []rune(s), []rune{}
	for _, r := range in {
		if r != space {
			out = append(out, r)
		}
	}
	return string(out)
}
