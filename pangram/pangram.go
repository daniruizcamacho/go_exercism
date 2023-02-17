package pangram

import "strings"

func IsPangram(s string) bool {
	s = strings.ToLower(s)
	for l := 'a'; l <= 'z'; l++ {
		if !strings.ContainsRune(s, l) {
			return false
		}
	}
	return true
}
