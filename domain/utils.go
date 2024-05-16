package domain

import (
	"errors"
	"fmt"
	"unicode"
)

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

// ValidateName checks the format of the name and ensures it meets the system's requirements.
// Names must be between 3 and 20 characters, contain only alphanumeric characters,
// and must not contain spaces. This method is used for naming users, folders, and files.
func ValidateName(name string) error {

	if len(name) < 3 || len(name) > 20 {
		return fmt.Errorf("'%s' must be between 3 and 20 characters", name)
	}

	for _, r := range name {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return fmt.Errorf("Error: The '%s' contain invalid chars", name)
		}
	}

	// Check for spaces in the username
	if containsSpace(name) {
		return errors.New("username must not contain spaces")
	}
	return nil
}
