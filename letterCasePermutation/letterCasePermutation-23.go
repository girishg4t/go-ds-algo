package lettercasepermutation

import (
	"unicode"
)

func LetterCasePermutation(s string) []string {
	upperStore := make(map[byte]byte)
	lowerStore := make(map[byte]byte)

	result := []string{s}
	for _, c := range s {
		upper := byte(unicode.ToUpper(c))
		lower := byte(unicode.ToLower(c))
		b := byte(c)
		if upper != b {
			upperStore[b] = upper
		}
		if lower != b {
			lowerStore[b] = lower
		}
	}

	for i, c := range s {
		b := byte(c)

		if val, ok := upperStore[b]; ok {
			result = append(result, s[:i]+string(val)+s[i+1:])
		}

		if val, ok := lowerStore[b]; ok {
			result = append(result, s[:i]+string(val)+s[i+1:])
		}
	}

	if len(upperStore) == 1 && len(lowerStore) == 1 {
		temp := ""
		for _, c := range s {
			b := byte(c)
			char := string(c)
			if val, ok := upperStore[b]; ok {
				char = string(val)
			}
			if val, ok := lowerStore[b]; ok {
				char = string(val)
			}
			temp = temp + char
		}
		result = append(result, temp)
	}

	if len(upperStore) > 1 {
		tempUpper := ""
		for _, c := range s {
			b := byte(c)
			if val, ok := upperStore[b]; ok {
				tempUpper = tempUpper + string(val)
			} else {
				tempUpper = tempUpper + string(c)
			}
		}
		if tempUpper != s {
			result = append(result, tempUpper)
		}
	}

	if len(lowerStore) > 1 {
		tempLower := ""
		for _, c := range s {
			b := byte(c)
			if val, ok := lowerStore[b]; ok {
				tempLower = tempLower + string(val)
			} else {
				tempLower = tempLower + string(c)
			}
		}

		if tempLower != s {
			result = append(result, tempLower)
		}
	}
	return result
}
