package domain

import (
	"fmt"
	"strings"
	"sync"
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
	if err := ValidateName(name); err != nil {
		return fmt.Errorf("Error: The '%s' contain invalid chars", name)
	}

	// Check if the username already exists in the map.
	if _, exists := u.list[name]; exists {
		return fmt.Errorf("Error: The '%s' has already existed", name)
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
		return nil, fmt.Errorf("Error: The '%s' doesn't exist", username)
	}

	return folders, nil
}
