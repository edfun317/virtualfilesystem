package domain

import "unicode"

// sortBy it's an enumerate used to specify the sorting type.
type sortBy int

const (
	sortByName sortBy = iota
	sortByCreated
)

// sortOrder it's an enumerate used to specify the sort order.
type sortOrder int

const (
	ascending sortOrder = iota
	descending
)

// containsSpace checks if the input string contains any space characters.
func containsSpace(s string) bool {
	for _, r := range s {
		if unicode.IsSpace(r) {
			return true
		}
	}
	return false
}
