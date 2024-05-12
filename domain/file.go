package domain

import (
	"sync"
	"time"
)

type (
	Files struct {
		List map[string]*File
		mux  sync.Mutex
	}
	File struct {
		Name        string
		Created     time.Time
		Description string
	}
)

// NewFiles creates a new files instance with initialized map.
func NewFiles() *Files {
	return &Files{
		List: make(map[string]*File),
	}
}

// AddFile adds a new file to the collection with the given name
func (f *Files) AddFile(name, description string) error {

	return nil
}

// RemoveFile removes a file from the collection by its name
func (f *Files) RemoveFile(name string) error {

	return nil
}

// ListFiles execute the GetSortedFiles and FilesFormatted methods to obtain sorted data
func (f *Files) ListFiles() ([]string, error) {

	return nil, nil
}

// GetSortedFiles returns a sorted slice of File based
func (f *Files) GetSortedFiles(by sortBy, order sortOrder) []File {

	return nil
}

// FilesFormatted formats a slice of File structures into a string slice for display purposes.
// The format includes the file name, optional description, creation date, and associated user and folder.
// Example output: "file1 this-is-file1 2023-01-01 15:00:20 folder1 user1"
func FilesFormatted(user, folder string, files []File) []string {

	return nil
}
