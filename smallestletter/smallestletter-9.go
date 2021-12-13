package smallestletter

func NextGreatestLetter(letters []byte, target byte) byte {
	for _, n := range letters {
		if n > target {
			return n
		}
	}
	return letters[0]
}
