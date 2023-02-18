package anagram

import "strings"

func Detect(subject string, candidates []string) []string {
	result := []string{}

	for _, candidate := range candidates {
		originalString := candidate
		candidate = strings.ToLower(candidate)
		subject = strings.ToLower(subject)
		if candidate == subject || len(subject) != len(candidate) {
			continue
		}

		charactersCounter := generateMapCounter(subject)
		for _, c := range candidate {
			charactersCounter[c] -= 1
		}

		isValid := true
		for _, counter := range charactersCounter {
			if counter != 0 {
				isValid = false
			}
		}

		if isValid {
			result = append(result, originalString)
		}
	}

	return result
}

func generateMapCounter(subject string) map[rune]int {
	charactersCounter := make(map[rune]int)

	for _, character := range subject {
		charactersCounter[character] += 1
	}

	return charactersCounter
}
