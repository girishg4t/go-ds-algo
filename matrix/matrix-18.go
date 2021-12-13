package matrix

func Matrix(matrix [][]int) {
	row := 0
	column := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				row = i
				column = j
			}
		}
	}

	i := row
	for i != -1 {
		matrix[i][column] = 0
		i--
	}
	i = row
	for i < len(matrix[row])-1 {
		matrix[i][column] = 0
		i++
	}

	i = column
	for i != -1 {
		matrix[row][i] = 0
		i--
	}
	i = column
	for i < len(matrix)-1 {
		matrix[row][i] = 0
		i++
	}
}
