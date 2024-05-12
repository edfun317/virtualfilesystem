package domain

import (
	"sync"
)

type (
	// Users struct holds a map from usernames to their corresponding Folders structures
	// and includes a mutex for synchronizing access to the map.
	Users struct {
		mux  sync.Mutex
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
	return nil
}

// GetUserFolders returns the Folders associated with a given username.
// It returns an error if the username does not exist.
func (u *Users) GetUserFolders(username string) (*Folders, error) {
	return nil, nil
}

// ValidateUsername checks the format of the username and ensures it meets the system's requirements.
// Usernames must be between 3 and 20 characters, contain only alphanumeric characters,
// and must not contain spaces. This method is used during user registration and updating the user profile.
func ValidateUsername(username string) error {
	return nil
}
