package codesignal

import (
	"strconv"
)

/* isLucky

Ticket numbers usually consist of an even number of digits. A ticket number is
considered lucky if the sum of the first half of the digits is equal to the
sum of the second half.

Given a ticket number n, determine if it's lucky or not.

Example

    For n = 1230, the output should be
    isLucky(n) = true;
    For n = 239017, the output should be
    isLucky(n) = false.

Input/Output

    [execution time limit] 4 seconds (go)

    [input] integer n

	A ticket number represented as a positive integer with an even number of
	digits.

    Guaranteed constraints:
    10 ≤ n < 106.

    [output] boolean
        true if n is a lucky ticket number, false otherwise.
*/

// Constant's are resolved at compile time.
const utf8zero = 48 // utf8one = 49, utf8two = 49...

func isLucky(n int) bool {
	// converting 'n' into a slice of UTF-8 codepoints
	runes := []rune(strconv.Itoa(n))
	rightSum, leftSum := 0, 0
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		// subtract the codepoints by the UTF-8 value for 0
		rightSum += int(runes[i]) - utf8zero
		leftSum += int(runes[j]) - utf8zero
	}
	return rightSum == leftSum
}
