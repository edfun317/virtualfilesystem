package event

import (
	"iscoollab/filesystem/domain"
	"testing"
)

var (
	userName = "testuser"
)

// setupUsers sets up a test environment with a user.
func setupUsers() *domain.Users {
	users := domain.NewUsers()
	users.AddUser(userName)
	return users
}

// setupFolder sets up a test environment with a user and a folder.
func setupFolder() *domain.Users {
	users := domain.NewUsers()
	users.AddUser(userName)
	folders, _ := users.GetUserFolders(userName)
	folders.AddFolder("testfolder", "", domain.NewFiles())
	return users
}

// setupFiles sets up a test environment with a user, a folder, and a file.
func setupFiles() *domain.Users {
	users := setupUsers()
	folders, _ := users.GetUserFolders(userName)
	folders.AddFolder("testfolder", "", domain.NewFiles())
	folder, _ := folders.FindFolder("testfolder")
	folder.TheFiles.AddFile("testfile", "description")
	return users
}

// TestCreateFile tests the CreateFile function.
func TestCreateFile(t *testing.T) {
	users := setupFolder()

	tests := []struct {
		name        string
		userName    string
		folderName  string
		fileName    string
		desc        string
		expectedErr string
	}{
		{"EmptyUserName", "", "testfolder", "testfile", "description", "Usage: create-file [username] [foldername] [filename] [description]?"},
		{"EmptyFolderName", "testuser", "", "testfile", "description", "Usage: create-file [username] [foldername] [filename] [description]?"},
		{"EmptyFileName", "testuser", "testfolder", "", "description", "Usage: create-file [username] [foldername] [filename] [description]?"},
		{"FolderNotFound", "testuser", "nonexistentfolder", "testfile", "description", "Error: The 'nonexistentfolder' doesn't exist"},
		{"ValidFileCreation", "testuser", "testfolder", "testfile", "description", ""},
		{"DuplicateFileName", "testuser", "testfolder", "testfile", "description", "Error: The 'testfile' has already existed."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateFile(users, tt.userName, tt.folderName, tt.fileName, tt.desc)
			if err != nil && err.Error() != tt.expectedErr {
				t.Errorf("expected error: %v, got: %v", tt.expectedErr, err)
			} else if err == nil && tt.expectedErr != "" {
				t.Errorf("expected error: %v, got: nil", tt.expectedErr)
			}
		})
	}
}

// TestCreateFolder tests the CreateFolder function.
func TestCreateFolder(t *testing.T) {
	users := setupUsers()

	tests := []struct {
		name        string
		userName    string
		folderName  string
		desc        string
		expectedErr string
	}{
		{"EmptyUserName", "", "testfolder", "description", "Usage: create-folder [username] [foldername] [description]?"},
		{"EmptyFolderName", userName, "", "description", "Usage: create-folder [username] [foldername] [description]?"},
		{"UserNotFound", "nonexistentuser", "testfolder", "description", "Error: The 'nonexistentuser' doesn't exist"},
		{"ValidFolderCreation", userName, "testfolder", "description", ""},
		{"DuplicateFolderName", userName, "testfolder", "description", "Error: The 'testfolder' has already existed."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateFolder(users, tt.userName, tt.folderName, tt.desc)
			if err != nil && err.Error() != tt.expectedErr {
				t.Errorf("expected error: %v, got: %v", tt.expectedErr, err)
			} else if err == nil && tt.expectedErr != "" {
				t.Errorf("expected error: %v, got: nil", tt.expectedErr)
			}
		})
	}
}

// TestDeleteFile tests the DeleteFile function.
func TestDeleteFile(t *testing.T) {
	users := setupFiles()

	tests := []struct {
		name        string
		userName    string
		folderName  string
		fileName    string
		expectedErr string
	}{
		{"EmptyUserName", "", "testfolder", "testfile", "Usage: delete-file [username] [foldername] [fileName]"},
		{"EmptyFolderName", "testuser", "", "testfile", "Usage: delete-file [username] [foldername] [fileName]"},
		{"EmptyFileName", "testuser", "testfolder", "", "Usage: delete-file [username] [foldername] [fileName]"},
		{"UserNotFound", "nonexistentuser", "testfolder", "testfile", "Error: The 'nonexistentuser' doesn't exist"},
		{"FolderNotFound", "testuser", "nonexistentfolder", "testfile", "Error: The 'nonexistentfolder' doesn't exist"},
		{"FileNotFound", "testuser", "testfolder", "nonexistentfile", "Error: The 'nonexistentfile' doesn't exist"},
		{"ValidFileDeletion", "testuser", "testfolder", "testfile", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := DeleteFile(users, tt.userName, tt.folderName, tt.fileName)
			if err != nil && err.Error() != tt.expectedErr {
				t.Errorf("expected error: %v, got: %v", tt.expectedErr, err)
			} else if err == nil && tt.expectedErr != "" {
				t.Errorf("expected error: %v, got: nil", tt.expectedErr)
			}
		})
	}
}

// TestDeleteFolder tests the DeleteFolder function.
func TestDeleteFolder(t *testing.T) {
	users := setupFiles()

	tests := []struct {
		name        string
		userName    string
		folderName  string
		expectedErr string
	}{
		{"EmptyUserName", "", "testfolder", "Usage: delete-folder [username] [foldername]"},
		{"EmptyFolderName", "testuser", "", "Usage: delete-folder [username] [foldername]"},
		{"UserNotFound", "nonexistentuser", "testfolder", "Error: The 'nonexistentuser' doesn't exist"},
		{"FolderNotFound", "testuser", "nonexistentfolder", "Error: The 'nonexistentfolder' doesn't exist"},
		{"ValidFolderDeletion", "testuser", "testfolder", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := DeleteFolder(users, tt.userName, tt.folderName)
			if err != nil && err.Error() != tt.expectedErr {
				t.Errorf("expected error: %v, got: %v", tt.expectedErr, err)
			} else if err == nil && tt.expectedErr != "" {
				t.Errorf("expected error: %v, got: nil", tt.expectedErr)
			}
		})
	}
}

// TestListFiles tests the ListFiles function.
func TestListFiles(t *testing.T) {
	users := setupFiles()

	tests := []struct {
		name        string
		userName    string
		folderName  string
		by          string
		order       string
		expectedErr string
	}{
		{"EmptyUserName", "", "testfolder", ByName, ASC, "Usage: list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]"},
		{"EmptyFolderName", "testuser", "", ByName, ASC, "Usage: list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]"},
		{"InvalidSortOption", "testuser", "testfolder", "invalid", ASC, "Usage: list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]"},
		{"InvalidSortOrder", "testuser", "testfolder", ByName, "invalid", "Usage: list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]"},
		{"FolderNotFound", "testuser", "nonexistentfolder", ByName, ASC, "Error: The 'nonexistentfolder' doesn't exist"},
		{"ValidListFiles", "testuser", "testfolder", ByName, ASC, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ListFiles(users, tt.userName, tt.folderName, tt.by, tt.order)
			if err != nil && err.Error() != tt.expectedErr {
				t.Errorf("expected error: %v, got: %v", tt.expectedErr, err)
			} else if err == nil && tt.expectedErr != "" {
				t.Errorf("expected error: %v, got: nil", tt.expectedErr)
			}
		})
	}
}

// TestListFolders tests the ListFolders function.
func TestListFolders(t *testing.T) {
	users := setupFiles()

	tests := []struct {
		name        string
		userName    string
		by          string
		order       string
		expectedErr string
	}{
		{"EmptyUserName", "", ByName, ASC, "Usage: list-folders [username] [--sort-name|--sort-created] [asc|desc]"},
		{"InvalidSortOption", "testuser", "invalid", ASC, "Usage: list-folders [username] [--sort-name|--sort-created] [asc|desc]"},
		{"InvalidSortOrder", "testuser", ByName, "invalid", "Usage: list-folders [username] [--sort-name|--sort-created] [asc|desc]"},
		{"UserNotFound", "nonexistentuser", ByName, ASC, "Error: The 'nonexistentuser' doesn't exist"},
		{"ValidListFolders", "testuser", ByName, ASC, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ListFolders(users, tt.userName, tt.by, tt.order)
			if err != nil && err.Error() != tt.expectedErr {
				t.Errorf("expected error: %v, got: %v", tt.expectedErr, err)
			} else if err == nil && tt.expectedErr != "" {
				t.Errorf("expected error: %v, got: nil", tt.expectedErr)
			}
		})
	}
}

// TestRegister tests the Register function.
func TestRegister(t *testing.T) {
	users := setupUsers()

	tests := []struct {
		name        string
		userName    string
		expectedErr string
	}{
		{"EmptyUserName", "", "Usage: register [username]"},
		{"ValidRegister", "newuser", ""},
		{"DuplicateRegister", "testuser", "Error: The 'testuser' has already existed"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Register(users, tt.userName)
			if err != nil && err.Error() != tt.expectedErr {
				t.Errorf("expected error: %v, got: %v", tt.expectedErr, err)
			} else if err == nil && tt.expectedErr != "" {
				t.Errorf("expected error: %v, got: nil", tt.expectedErr)
			}
		})
	}
}

// TestRenameFolder tests the RenameFolder function.
func TestRenameFolder(t *testing.T) {
	users := setupFiles()

	tests := []struct {
		name        string
		userName    string
		folderName  string
		newFolder   string
		expectedErr string
	}{
		{"EmptyUserName", "", "testfolder", "newfolder", "Usage: rename-folder [username] [foldername] [new-folder-name]"},
		{"EmptyFolderName", "testuser", "", "newfolder", "Usage: rename-folder [username] [foldername] [new-folder-name]"},
		{"EmptyNewFolderName", "testuser", "testfolder", "", "Usage: rename-folder [username] [foldername] [new-folder-name]"},
		{"UserNotFound", "nonexistentuser", "testfolder", "newfolder", "Error: The 'nonexistentuser' doesn't exist"},
		{"FolderNotFound", "testuser", "nonexistentfolder", "newfolder", "Error: The 'nonexistentfolder' doesn't exist"},
		{"ValidRename", "testuser", "testfolder", "newfolder", ""},
		{"DuplicateFolderName", "testuser", "newfolder", "newfolder", "Error: The 'newfolder' has already existed"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := RenameFolder(users, tt.userName, tt.folderName, tt.newFolder)
			if err != nil && err.Error() != tt.expectedErr {
				t.Errorf("expected error: %v, got: %v", tt.expectedErr, err)
			} else if err == nil && tt.expectedErr != "" {
				t.Errorf("expected error: %v, got: nil", tt.expectedErr)
			}
		})
	}
}
