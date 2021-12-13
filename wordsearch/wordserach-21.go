package wordsearch

func Exist(board [][]byte, word string) bool {

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if string(word[0]) == string(board[i][j]) {
				if checkSurrounding(board, word[1:], i, j) {
					return true
				}
			}
		}
	}

	return false
}

func checkSurrounding(board [][]byte, word string, x, y int) bool {
	for i, c := range word {
		found := false
		up := x - 1
		down := x + 1
		left := y - 1
		right := y + 1

		if up > -1 && string(c) == string(board[x-1][y]) {
			found = true
			checkSurrounding(board, word[i+1:], x-1, y)
		}

		if down < len(board) && string(c) == string(board[x+1][y]) {
			found = true
			checkSurrounding(board, word[i+1:], x+1, y)
		}

		if right < len(board[0]) && string(c) == string(board[x][y+1]) {
			found = true
			checkSurrounding(board, word[i+1:], x, y+1)
		}
		if left > -1 && string(c) == string(board[x][y-1]) {
			found = true
			checkSurrounding(board, word[i+1:], x, y-1)
		}
		if !found {
			return false
		}
	}
	return true
}
