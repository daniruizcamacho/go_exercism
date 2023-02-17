package pangram

import "unicode"

func IsPangram(input string) bool {
	alphabetMap := make(map[string]bool)
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for _, letter := range alphabet {
		alphabetMap[string(letter)] = false
	}

	for _, v := range input {
		alphabetMap[string(unicode.ToLower(v))] = true
	}

	for _, v := range alphabetMap {
		if !v {
			return false
		}
	}

	return true
}
