// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(number int) bool {
	if number%4 == 0 && number%100 != 0 {
		return true
	}

	if number%4 == 0 && number%100 == 0 && number%400 == 0 {
		return true
	}

	return false
}
