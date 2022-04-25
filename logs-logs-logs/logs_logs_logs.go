package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	applications := map[rune]string{
		'❗': "recommendation",
		'🔍': "search",
		'☀': "weather",
	}
	for _, character := range log {
		if val, exists := applications[character]; exists {
			return val
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var newString string
	for _, v := range log {
		if v == oldRune {
			newString += string(newRune)
			continue
		}

		newString += string(v)
	}

	return newString
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
