package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	freq := make(map[string]int)
	phrase = strings.ToLower(phrase)
	wordPattern := regexp.MustCompile("[a-z0-9]+(['][a-z0-9]+)?")
	for _, word := range wordPattern.FindAllString(phrase, -1) {
		freq[word] += 1
	}

	return freq
}
