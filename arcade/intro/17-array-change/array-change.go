package codesignal

/* arrayChange
You are given an array of integers. On each move you are allowed to increase
exactly one of its element by one. Find the minimal number of moves required to
obtain a strictly increasing sequence from the input.

Example

For inputArray = [1, 1, 1], the output should be
arrayChange(inputArray) = 3.

Input/Output

    [execution time limit] 4 seconds (go)

    [input] array.integer inputArray

    Guaranteed constraints:
    3 ≤ inputArray.length ≤ 105,
    -105 ≤ inputArray[i] ≤ 105.

    [output] integer
		The minimal number of moves needed to obtain a strictly increasing
		sequence from inputArray.

		It's guaranteed that for the given test cases the answer always fits
		signed 32-bit integer type.

*/
func arrayChange(ints []int) int {
	steps := 0
	for i, length := 1, len(ints); i < length; i++ { // caching len(var) saves CPU cycles.
		prev, current := ints[i-1], ints[i]
		if current <= prev {
			increase := prev + 1
			ints[i] = increase
			steps += increase - current
		}
	}
	return steps
}
