package domain

import (
	"testing"
	"time"
)

func TestNewFiles(t *testing.T) {
	files := NewFiles()
	if files == nil {
		t.Fatal("NewFiles() should not return nil")
	}
	if files.List == nil {
		t.Errorf("New Files list should be init")
	}
}

// TestAddFile
// step:
// 1. make sure adding a file doesn't mess up.
// 2. check if the file show up in the list after adding.
// 3. kicks up a fuss when try to add the same file twice.
func TestAddFile(t *testing.T) {

	f := NewFiles()
	name := "testFile"

	err := f.AddFile(name, "A test file")
	if err != nil {
		t.Errorf("Failed to add file: %v", err)
	}

	if _, exists := f.List[name]; !exists {
		t.Errorf("File '%s' was not added correctly", name)
	}

	err = f.AddFile("testFile", "A test file")
	if err == nil {
		t.Error("Expected error when adding a duplicate file, got nil")
	}
}

// TestRmoveFile
// Step:
// 1. add the file to the list
// 2. make sure remove a file doesn't mess up.
// 3. check it the file has been removed from the list
// 4. kick up a fuss when try to remove the same file twice.
func TestRemoveFile(t *testing.T) {

	f := NewFiles()
	name := "testFile"

	f.AddFile("testFile", "")

	err := f.RemoveFile(name)
	if err != nil {
		t.Errorf("Failed to remove file: %v", err)
	}

	if _, exists := f.List[name]; exists {
		t.Errorf("Folder '%s' was not removed correctly", name)
	}

	// Try to remove the same file again
	err = f.RemoveFile(name)
	if err == nil {
		t.Error("Expected error when trying to remove a non-existent file, got nil")
	}
}

func setupFiles() *Files {
	f := NewFiles()
	f.List["file1"] = &File{Name: "file1", Created: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC), Description: "First file"}
	f.List["file2"] = &File{Name: "file2", Created: time.Date(2023, 1, 2, 12, 0, 0, 0, time.UTC), Description: "Second file"}
	f.List["file3"] = &File{Name: "file3", Created: time.Date(2023, 1, 3, 12, 0, 0, 0, time.UTC), Description: "Third file"}
	return f
}

// TestGetSortedFilesByName
// case: sort by name, and sorting by asc or desc
func TestGetSortedFilesByName(t *testing.T) {
	f := setupFiles()

	// Test ascending order by name
	sortedFiles := f.GetSortedFiles(sortByName, ascending)
	if sortedFiles[0].Name > sortedFiles[1].Name {
		t.Errorf("Expected files to be sorted by name in ascending order, got %v before %v", sortedFiles[0].Name, sortedFiles[1].Name)
	}

	// Test descending order by name
	sortedFiles = f.GetSortedFiles(sortByName, descending)
	if sortedFiles[0].Name < sortedFiles[1].Name {
		t.Errorf("Expected files to be sorted by name in descending order, got %v before %v", sortedFiles[0].Name, sortedFiles[1].Name)
	}
}

// TestGetSortedFilesByName
// case: sort by created, and sorting by asc or desc
func TestGetSortedFilesByCreated(t *testing.T) {
	f := setupFiles()

	// Test ascending order by created date
	sortedFiles := f.GetSortedFiles(sortByCreated, ascending)
	if sortedFiles[0].Created.After(sortedFiles[1].Created) {
		t.Errorf("Expected files to be sorted by created date in ascending order, got %v before %v", sortedFiles[0].Created, sortedFiles[1].Created)
	}

	// Test descending order by created date
	sortedFiles = f.GetSortedFiles(sortByCreated, descending)
	if sortedFiles[0].Created.Before(sortedFiles[1].Created) {
		t.Errorf("Expected files to be sorted by created date in descending order, got %v before %v", sortedFiles[0].Created, sortedFiles[1].Created)
	}
}

// TestFilesFormatted
// case: check the result is equal to the expected
func TestFilesFormatted(t *testing.T) {
	files := []File{
		{Name: "file1", Description: "this-is-file1", Created: time.Date(2023, 1, 1, 15, 0, 0, 0, time.UTC)},
		{Name: "file2", Created: time.Date(2023, 1, 1, 15, 0, 2, 0, time.UTC)},
	}
	user := "user1"
	folder := "folder1"
	expected := []string{"file1 this-is-file1 2023-01-01 15:00:00 folder1 user1", "file2 2023-01-01 15:00:02 folder1 user1"}

	result := FilesFormatted(user, folder, files)
	if len(result) != len(expected) {
		t.Fatalf("Expected %d formatted strings, got %d", len(expected), len(result))
	}

	for i, str := range result {
		if str != expected[i] {
			t.Errorf("Expected '%s', got '%s'", expected[i], str)
		}
	}
}

// TestListFiles
// Step:
// 1.setup a Files instance with some files
// 2.setup the expected slice string
// 3.test for ascending sort by creation
// 4.test for descending sort by names
// 5.test empty list handling
func TestListFiles(t *testing.T) {
	f := &Files{
		List: map[string]*File{
			"file1": {Name: "file1", Created: time.Date(2023, 1, 1, 14, 0, 0, 0, time.UTC)}, // 1 day ago
			"file2": {Name: "file2", Created: time.Date(2023, 1, 1, 15, 0, 0, 0, time.UTC)}, // now
		},
	}

	expected := []string{"file1 2023-01-01 14:00:00 folder1 user1", "file2 2023-01-01 15:00:00 folder1 user1"}
	user := "user1"
	folder := "folder1"

	// test for ascending sort by creation
	files, err := f.ListFiles(user, folder, ByCreated, ASC)
	if err != nil {
		t.Errorf("ListFiles returned an error: %v", err)
	}

	if len(files) != 2 || files[0] != expected[0] {
		t.Errorf("Expected ['file1', 'file2'], got %v", files)
	}

	//test for descending sort by names
	files, err = f.ListFiles(user, folder, ByName, DESC)
	if err != nil {
		t.Errorf("ListFiles returned an error: %v", err)
	}
	if len(files) != 2 || files[0] != expected[1] {
		t.Errorf("Expected ['file2', 'file1'], got %v", files)
	}

	// Test empty list handling
	f = &Files{List: map[string]*File{}}
	files, err = f.ListFiles("john", "emptyFolder", "", "")
	if err == nil {
		t.Error("Expected ListFiles to return a warning when the folder is empty.")
	}
	if len(files) != 0 {
		t.Errorf("Expected empty list, got %v", files)
	}
}
