package domain

import "unicode"

// containsSpace checks if the input string contains any space characters.
func containsSpace(s string) bool {
	for _, r := range s {
		if unicode.IsSpace(r) {
			return true
		}
	}
	return false
}
