package codesignal

/*
Given an array of integers, find the pair of adjacent elements that has the
largest product and return that product.

Example

For inputArray = [3, 6, -2, -5, 7, 3], the output should be
adjacentElementsProduct(inputArray) = 21.

7 and 3 produce the largest product.

Input/Output

    [execution time limit] 4 seconds (go)

    [input] array.integer inputArray

    An array of integers containing at least two elements.

    Guaranteed constraints:
    2 ≤ inputArray.length ≤ 10,
    -1000 ≤ inputArray[i] ≤ 1000.

    [output] integer
		The largest product of adjacent elements.
*/

func adjacentElementsProduct(input []int) (largest int) {
	largest = input[0] * input[1]
	for i, length := 2, len(input); i < length; i++ {
		product := input[i-1] * input[i]
		if product > largest {
			largest = product
		}
	}
	return largest
}
