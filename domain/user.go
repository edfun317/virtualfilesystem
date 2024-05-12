package domain

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"unicode"
)

type (
	// Users struct holds a map from usernames to their corresponding Folders structures
	// and includes a mutex for synchronizing access to the map.
	Users struct {
		mu   sync.Mutex
		list map[string]*Folders
	}
)

// NewUsers initializes and returns a new instance of a Users map.
func NewUsers() *Users {
	return &Users{
		list: make(map[string]*Folders),
	}
}

// AddUser adds a new user to the Users map after validating the username.
// It ensures usernames are unique and meet the defined format requirements.
func (u *Users) AddUser(username string) error {
	// Normalize username to lowercase to ensure case-insensitivity.
	name := strings.ToLower(username)

	// Validate the username format.
	if err := ValidateUsername(name); err != nil {
		return fmt.Errorf("Error:The '%s' contain invalid chars", name)
	}

	// Check if the username already exists in the map.
	if _, exists := u.list[name]; exists {
		return fmt.Errorf("Error:The '%s' has already existed", name)
	}

	// Create a new Folders instance for the new user.
	u.list[name] = NewFolders()
	return nil
}

// GetUserFolders returns the Folders associated with a given username.
// It returns an error if the username does not exist.
func (u *Users) GetUserFolders(username string) (*Folders, error) {
	folders, exists := u.list[username]
	if !exists {
		return nil, fmt.Errorf("user '%s' does not exist", username)
	}

	return folders, nil
}

// ValidateUsername checks the format of the username and ensures it meets the system's requirements.
// Usernames must be between 3 and 20 characters, contain only alphanumeric characters,
// and must not contain spaces. This method is used during user registration and updating the user profile.
func ValidateUsername(username string) error {

	if len(username) < 3 || len(username) > 20 {
		return errors.New("username must be between 3 and 20 characters")
	}

	for _, r := range username {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return errors.New("username must contain only alphanumeric characters")
		}
	}

	// Check for spaces in the username
	if containsSpace(username) {
		return errors.New("username must not contain spaces")
	}
	return nil
}
