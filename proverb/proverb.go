// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	proverb := make([]string, 0)

	if nil == rhyme || 0 == len(rhyme) {
		return proverb
	}

	prev := ""
	first := ""
	for _, value := range rhyme {
		if prev == "" {
			prev = value
			first = value
			continue
		}

		proverb = append(proverb, fmt.Sprintf("For want of a %s the %s was lost.", prev, value))
		prev = value
	}

	proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", first))

	return proverb
}
