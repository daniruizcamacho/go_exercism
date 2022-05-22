package isogram

import "strings"

func IsIsogram(word string) bool {
	lettersExists := make(map[string]bool)

	for _, v := range word {
		letter := strings.ToLower(string(v))

		if letter == " " || letter == "-" {
			continue
		}

		if _, ok := lettersExists[letter]; ok {
			return false
		}

		lettersExists[letter] = true
	}

	return true
}
