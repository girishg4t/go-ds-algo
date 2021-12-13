package spiralmatrix

func SpiralOrder(matrix [][]int) []int {
	maxRow := len(matrix) - 1
	maxCol := len(matrix[0]) - 1

	result := []int{}
	dir := 0
	startRow := 0
	startColumn := 0
	for startRow <= maxRow && startColumn <= maxCol {
		switch dir {
		// right
		case 0:
			for i := startRow; i <= maxCol; i++ {
				result = append(result, matrix[startRow][i])
			}
			dir = 1
			startRow++
		// down
		case 1:
			for i := startRow; i <= maxRow; i++ {
				result = append(result, matrix[i][maxCol])
			}

			dir = 2
			maxCol--
		//left
		case 2:
			for i := maxCol; i >= startColumn; i-- {
				result = append(result, matrix[maxRow][i])
			}
			dir = 3
			maxRow--
		// up
		case 3:
			for i := maxRow; i >= startRow; i-- {
				result = append(result, matrix[i][startColumn])
			}
			dir = 0
			startColumn++
		}
	}

	return result
}
