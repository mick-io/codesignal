package codesignal

/* minesweeper

In the popular Minesweeper game you have a board with some mines and those cells
that don't contain a mine have a number in it that indicates the total number of
mines in the neighboring cells. Starting off with some arrangement of mines we
want to create a Minesweeper game setup.

Example

For

matrix = [[true, false, false],
          [false, true, false],
          [false, false, false]]

the output should be

minesweeper(matrix) = [[1, 2, 1],
                       [2, 1, 1],
                       [1, 1, 1]]

Check out the image below for better understanding:

Input/Output

    [execution time limit] 4 seconds (go)

    [input] array.array.boolean matrix

	A non-empty rectangular matrix consisting of boolean values - true if the
	corresponding cell contains a mine, false otherwise.

    Guaranteed constraints:
    2 ≤ matrix.length ≤ 100,
    2 ≤ matrix[0].length ≤ 100.

    [output] array.array.integer
		Rectangular matrix of the same size as matrix each cell of which contains
		an integer equal to the number of mines in the neighboring cells. Two
		cells are called neighboring if they share at least one corner.

*/

func minesweeper(minefield [][]bool) (mineMatrix [][]int) {
	nRows, nColumn := len(minefield), len(minefield[0])
	minefield = buffer(minefield)
	for x := 1; x <= nRows; x++ {
		above, adjacent, below := minefield[x-1], minefield[x], minefield[x+1]
		detections := make([]int, nColumn)
		for y := 1; y <= nColumn; y++ {
			nMines := countMines(above[y-1:y+2], adjacent[y-1:y+2], below[y-1:y+2])
			if adjacent[y] {
				nMines--
			}
			detections[y-1] = nMines
		}
		mineMatrix = append(mineMatrix, detections)
	}
	return
}

func countMines(slices ...[]bool) (nMines int) {
	for _, slice := range slices {
		for _, b := range slice {
			if b {
				nMines++
			}
		}
	}
	return
}

// Buffering the minefield with false to avoid excessive conditionals
func buffer(matrix [][]bool) [][]bool {
	for i, bools := range matrix {
		bools = append([]bool{false}, bools...)
		matrix[i] = append(bools, false)
	}
	length := len(matrix[0])
	matrix = append([][]bool{make([]bool, length)}, matrix...)
	return append(matrix, make([]bool, length))
}
