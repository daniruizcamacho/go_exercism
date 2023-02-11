package etl

import "strings"

func Transform(input map[int][]string) map[string]int {
	output := make(map[string]int)
	for points, letters := range input {
		for _, letter := range letters {
			transformedLetter := strings.ToLower(letter)
			output[transformedLetter] = points
		}
	}
	return output
}
