package event

import (
	"iscoollab/filesystem/domain"
	"testing"
)

func TestCreateFolder(t *testing.T) {
	// Setup Users and Folders
	users := domain.NewUsers()
	username := "testuser"
	users.AddUser(username)

	// Successful case: Adding a new folder
	err := CreateFolder(users, username, "projects", "Project folder")
	if err != nil {
		t.Errorf("CreateFolder failed when it should succeed: %v", err)
	}

	// Error case: Trying to add a duplicate folder
	err = CreateFolder(users, username, "projects", "Project folder")
	if err == nil {
		t.Errorf("CreateFolder should fail on duplicate folder, but it did not")
	}
}
