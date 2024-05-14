package domain

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFilesManager struct {
	mock.Mock
}

func (m *MockFilesManager) AddFile(name, desc string) error {
	args := m.Called(name, desc)
	return args.Error(0)
}

func (m *MockFilesManager) RemoveFile(name string) error {
	args := m.Called(name)
	return args.Error(0)
}

func (m *MockFilesManager) ListFiles(user, folder, by, order string) ([]string, error) {
	args := m.Called(user, folder, by, order)
	return args.Get(0).([]string), args.Error(1)
}

func TestNewFolders(t *testing.T) {
	folders := NewFolders()
	if folders == nil {
		t.Fatal("NewFolders() returned nil")
	}
	if folders.List == nil {
		t.Errorf("New folder list should be init")
	}
}

func TestAddFolder(t *testing.T) {
	filesManager := new(MockFilesManager) // Create a new mock FilesManager
	folders := NewFolders()
	name := "testFolder"
	// Case 1: Valid folder addition
	err := folders.AddFolder(name, "", filesManager)
	if err != nil {
		t.Errorf("Failed to add folder: %v", err)
	}

	// Verify folder is added
	if _, exists := folders.List[name]; !exists {
		t.Errorf("Folder '%s' was not added correctly", name)
	}

	// Case 2: Attempt to add a duplicate folder
	err = folders.AddFolder(name, "", filesManager)
	if err == nil || !strings.Contains(err.Error(), "already existed") {
		t.Errorf("Expected error for duplicate folder, but got none or the wrong error: %v", err)
	}

	// Case 3: Adding a folder with invalid characters
	err = folders.AddFolder("Invalid/Name", "Invalid name", filesManager)
	if err == nil || !strings.Contains(err.Error(), "invalid chars") {
		t.Errorf("Expected validation error for folder name, but got none or the wrong error: %v", err)
	}
}

func TestFindFolder(t *testing.T) {

	filesManager := new(MockFilesManager) // Create a new mock FilesManager
	f := NewFolders()
	name := "testFolder"

	// Test adding a folder
	err := f.AddFolder(name, "", filesManager)
	if err != nil {
		t.Errorf("Failed to add folder: %v", err)
	}

	_, err = f.FindFolder(name)
	if err != nil {

		t.Errorf("Failed to add folder: %v", err)
	}
}
func TestRemoveFolder(t *testing.T) {

	filesManager := new(MockFilesManager) // Create a new mock FilesManager
	f := NewFolders()
	name := "testFolder"

	// Test adding a folder
	err := f.AddFolder(name, "", filesManager)
	if err != nil {
		t.Errorf("Failed to add folder: %v", err)
	}

	// Test removing the folder
	err = f.RemoveFolder(name)
	if err != nil {
		t.Errorf("Failed to remove folder: %v", err)
	}

	if _, exists := f.List[name]; exists {
		t.Errorf("Folder '%s' was not removed correctly", name)
	}

	// Try to remove the same folder again
	err = f.RemoveFolder(name)
	if err == nil {
		t.Error("Expected error when trying to remove a non-existent folder, got nil")
	}
}

func createTestFolders() *Folders {
	f := NewFolders()
	f.List["folderA"] = &Folder{Name: "folderA", Created: time.Now().Add(-time.Hour), Description: "Desc A"}
	f.List["folderB"] = &Folder{Name: "folderB", Created: time.Now(), Description: "Desc B"}
	return f
}

func TestGetSortedFoldersByName(t *testing.T) {
	f := createTestFolders()

	// Test ascending order by name
	sorted := f.GetSortedFolders(sortByName, ascending)
	assert.Equal(t, "folderA", sorted[0].Name, "Expected folderA to be first in ascending order by name")
	assert.Equal(t, "folderB", sorted[1].Name, "Expected folderB to be second in ascending order by name")

	// Test descending order by name
	sorted = f.GetSortedFolders(sortByName, descending)
	assert.Equal(t, "folderB", sorted[0].Name, "Expected folderB to be first in descending order by name")
	assert.Equal(t, "folderA", sorted[1].Name, "Expected folderA to be second in descending order by name")
}

func TestGetSortedFoldersByCreated(t *testing.T) {
	f := createTestFolders()

	// Test ascending order by created date
	sorted := f.GetSortedFolders(sortByCreated, ascending)
	assert.True(t, sorted[0].Created.Before(sorted[1].Created), "Expected older folder to be first in ascending order by created date")

	// Test descending order by created date
	sorted = f.GetSortedFolders(sortByCreated, descending)
	assert.True(t, sorted[0].Created.After(sorted[1].Created), "Expected newer folder to be first in descending order by created date")
}

func TestRenameFolder(t *testing.T) {

	filesManager := new(MockFilesManager) // Create a new mock FilesManager
	folders := NewFolders()
	originalName := "originalFolder"
	newName := "newFolder"

	// Setup initial folder
	folders.AddFolder(originalName, "", filesManager)

	// Test renaming to a new name
	err := folders.Rename(originalName, newName)
	if err != nil {
		t.Errorf("Failed to rename folder: %v", err)
	}
	if _, ok := folders.List[newName]; !ok {
		t.Errorf("Folder was not renamed in the list. Expected '%s' to exist", newName)
	}
	if _, ok := folders.List[originalName]; ok {
		t.Errorf("Original folder still exists after renaming. Expected '%s' to be deleted", originalName)
	}

	// Test renaming to an already existing name
	folders.AddFolder(originalName, "", filesManager) // Re-add original folder for further testing
	err = folders.Rename(originalName, newName)
	if err == nil {
		t.Error("Expected error when renaming to an existing folder name, but got nil")
	}

	// Test renaming a non-existing folder
	err = folders.Rename("nonExisting", "shouldFail")
	if err == nil {
		t.Error("Expected error when trying to rename a non-existing folder, but got nil")
	}
}

// TestListFolders
// Step:
// 1.setup a Folders instance with some files
// 2.setup the expected slice string
// 3.test for ascending sort by creation
// 4.test for descending sort by names
// 5.test empty list handling
func TestListFolders(t *testing.T) {

	f := &Folders{
		List: map[string]*Folder{
			"folder1": {Name: "folder1", Created: time.Date(2023, 1, 1, 14, 0, 0, 0, time.UTC)}, // 1 day ago
			"folder2": {Name: "folder2", Created: time.Date(2023, 1, 1, 15, 0, 0, 0, time.UTC)}, // now
		},
	}

	expected := []string{"folder1 2023-01-01 14:00:00 user1", "folder2 2023-01-01 15:00:00 user1"}
	user := "user1"

	// test for ascending sort by creation
	files, err := f.ListFolders(user, "sort-created", "asc")
	if err != nil {
		t.Errorf("ListFolders returned an error: %v", err)
	}

	if len(files) != 2 || files[0] != expected[0] {
		t.Errorf("Expected ['file1', 'file2'], got %v", files)
	}

	//test for descending sort by names
	files, err = f.ListFolders(user, "", "desc")
	if err != nil {
		t.Errorf("ListFolders returned an error: %v", err)
	}
	if len(files) != 2 || files[0] != expected[1] {
		t.Errorf("Expected ['file2', 'file1'], got %v", files)
	}

	// Test empty list handling
	f = &Folders{List: map[string]*Folder{}}
	files, err = f.ListFolders("john", "sort-created", "asc")
	if err != nil {
		t.Errorf("ListFolders returned an error with empty list: %v", err)
	}
	if len(files) != 0 {
		t.Errorf("Expected empty list, got %v", files)
	}
}

func TestFoldersFormatted(t *testing.T) {
	// test data
	user := "user1"
	folders := []Folder{
		{
			Name:        "folder1",
			Description: "this-is-folder-1",
			Created:     time.Date(2024, 1, 1, 15, 0, 0, 0, time.UTC),
		},
		{
			Name:    "folder2",
			Created: time.Date(2024, 1, 1, 15, 0, 10, 0, time.UTC),
		},
	}

	// Expected output
	expected := []string{
		"folder1 this-is-folder-1 2024-01-01 15:00:00 user1",
		"folder2 2024-01-01 15:00:10 user1",
	}

	result := FoldersFormatted(user, folders)

	// Check the length of the result
	if len(result) != len(expected) {
		t.Errorf("Expected result length %d, got %d", len(expected), len(result))
	}

	// Compare each element
	for i, res := range result {
		if res != expected[i] {
			t.Errorf("Expected result at index %d to be '%s', got '%s'", i, expected[i], res)
		}
	}
}
