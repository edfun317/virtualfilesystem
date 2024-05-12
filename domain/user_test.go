package domain

import (
	"testing"
)

func TestNewUsers(t *testing.T) {
	users := NewUsers()
	if len(users.list) != 0 {
		t.Errorf("Expected empty Users map, got %v", users)
	}
}

func TestAddUser(t *testing.T) {
	users := NewUsers()
	err := users.AddUser("User123")
	if err != nil {
		t.Errorf("Failed to add user: %v", err)
	}

	// Test adding a user with invalid characters
	err = users.AddUser("Invalid User!")
	if err == nil {
		t.Error("Expected error when adding a user with invalid characters, got nil")
	}

	// Test adding the same user again
	err = users.AddUser("user123") // Should fail due to case insensitivity
	if err == nil {
		t.Error("Expected error when adding a duplicate user, got nil")
	}
}

func TestGetUserFolders(t *testing.T) {

	users := NewUsers()
	users.AddUser("User123")

	_, err := users.GetUserFolders("user123") // Case insensitive check
	if err != nil {
		t.Errorf("Failed to retrieve folders for existing user: %v", err)
	}

	_, err = users.GetUserFolders("nonexistent")
	if err == nil {
		t.Error("Expected error when retrieving folders for a nonexistent user, got nil")
	}
}
