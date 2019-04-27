package codesignal

import (
	"strconv"
)

/*
When subtracting integers by hand, you can use a strategy called regrouping.
(This strategy is also sometimes called borrowing.) In this strategy,
corresponding digits are subtracted from each other one by one, from right to
left, unless the digit of the minuend (the number being subtracted from) is
smaller than the corresponding digit of the subtrahend (the number being
subtracted). In this case, you have to borrow 10 from the digit of the minuend
immediately to the left of the current one. You can see an example of this
method at work here.

When you subtract using regrouping, all the numbers that are being subtracted
(the subtrahend) remain the same, while the numbers that are being subtracted
from (the minuend) may change. Given two integers of the same length, return an
array containing the numbers that are being subtracted from when performing
subtraction by hand, from right to left.

Example

For minuend = 2024 and subtrahend = 1234, the output should be
subtractionByRegrouping(minuend, subtrahend) = [4, 12, 9, 1].

When subtracting 1234 from 2024, we actually consider the following pairs
(going from right to left):

    4 and 4;
    12 (10 was borrowed from 0, temporarily making it equal to -1, & added to 2) & 3;
    9 (again, 10 was borrowed from 2) and 2;
    1 and 1.

Input/Output

    [execution time limit] 4 seconds (go)

    [input] integer minuend

    The number that the subtrahend is being subtracted from.

    Guaranteed constraints:
    1 ≤ minuend ≤ 109.

    [input] integer subtrahend

    The number being subtracted from the minuend.

    subtrahend is guaranteed to have the same number of digits as minuend.

    Guaranteed constraints:
    1 ≤ subtrahend ≤ minuend.

    [output] array.integer
		An array containing the numbers that are being subtracted from when
		subtracting the subtrahend from the minuend using the regrouping
		strategy.
*/

// Constant's are resolved at compile time.
const utf8zero = 48 // utf8one = 49, utf8two = 49...

func subtractionByRegrouping(minuend int, subtrahend int) []int {
	base, subtract := splitInt(minuend), splitInt(subtrahend)
	output := make([]int, 0)
	for i := len(base) - 1; i >= 0; i-- {
		if base[i] < subtract[i] && i != 0 {
			for j := i - 1; j >= 0; j-- {
				if base[j] == 0 {
					continue
				}
				base[j]--
				for k := j + 1; k < i; k++ {
					base[k] += 9
				}
				base[i] += 10
				break
			}
		}
		output = append(output, base[i])
	}
	return output
}

func splitInt(n int) []int {
	in := []rune(strconv.Itoa(n))
	out := make([]int, len(in))
	for i, n := range in {
		out[i] = int(n) - utf8zero
	}
	return out
}
