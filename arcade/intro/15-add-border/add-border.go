package codesignal

/* addBorder
Given a rectangular matrix of characters, add a border of asterisks(*) to it.

Example

For

picture = ["abc",
           "ded"]

the output should be

addBorder(picture) = ["*****",
                      "*abc*",
                      "*ded*",
                      "*****"]

Input/Output

    [execution time limit] 4 seconds (go)

    [input] array.string picture

    A non-empty array of non-empty equal-length strings.

    Guaranteed constraints:
    1 ≤ picture.length ≤ 100,
    1 ≤ picture[i].length ≤ 100.

    [output] array.string
		The same matrix of characters, framed with a border of asterisks of
		width 1.

*/

// constants are resolved at compile time. They increase readability without hindering performance
const asterisk = "*"

func addBorder(picture []string) []string {
	for i, s := range picture {
		picture[i] = asterisk + s + asterisk
	}
	var border string
	for i, size := 0, len(picture[0]); i < size; i++ { // caching len(var) saves CPU cycles.
		border += asterisk
	}
	picture = append([]string{border}, picture...)
	return append(picture, border)
}
