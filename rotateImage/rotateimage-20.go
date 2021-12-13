package rotateimage

func Rotate(matrix [][]int) {
	maxRow := len(matrix) - 1
	maxCol := len(matrix[0]) - 1

	startRow := 0
	startColumn := 0

	// reverse
	for startRow <= maxRow {
		for startColumn <= maxCol {
			temp := matrix[startRow][startColumn]
			matrix[startRow][startColumn] = matrix[maxRow][startColumn]
			matrix[maxRow][startColumn] = temp
			startColumn++
		}
		startRow++
		maxRow--
		startColumn = 0
	}

	startRow = 0
	startColumn = 0
	//swap row and column

	row := startRow + 1
	column := startColumn + 1

	maxRow = len(matrix) - 1
	maxCol = len(matrix[0]) - 1
	for startRow <= maxRow && startColumn <= maxCol {
		for row <= maxRow && column <= maxCol {
			temp := matrix[startRow][column]
			matrix[startRow][column] = matrix[column][startRow]
			matrix[column][startRow] = temp
			column++
			row++
		}
		startRow++
		startColumn++
		row = startRow + 1
		column = startColumn + 1
	}

}
