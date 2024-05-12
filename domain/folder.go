package domain

import (
	"sync"
	"time"
)

type (
	// Folders contains a map of Folder structures indexed by their name.
	Folders struct {
		List map[string]*Folder
		mux  sync.Mutex
	}

	// Folder represents a directory in the virtual file system,
	// containing a name, creation time, the files and description.
	Folder struct {
		Name        string
		Created     time.Time
		TheFiles    TheFiles
		Description string
	}

	// TheFiles interface defines the methods to manage files within a folder.
	TheFiles interface {
		AddFile(name, desc string) error
		RemoveFile(name string) error
		ListFiles() ([]string, error)
	}
)

// NewFolders creates a new Folders instance with initialized map.
func NewFolders() *Folders {

	return &Folders{
		List: make(map[string]*Folder),
	}
}

// AddFolder adds a new folder to the collection with the given name.
func (f *Folders) AddFolder(name, description string) error {

	return nil
}

// RemoveFolder removes a folder from the collection by its name.
func (f *Folders) RemoveFolder(name string) error {

	return nil
}

// Rename rename a folder,
func (f *Folders) Rename(name, newName string) error {

	return nil
}

// GetSortedFolders returns a sorted slice of Folder based on the provided sort type and order.
func (f *Folders) GetSortedFolders(by sortBy, order sortOrder) []Folder {

	return nil
}

// FoldersFormatted formats a slice of Folder structures into a string slice
// for display purposes. The format includes the folder name, optional description,
// creation date, and associated user.
// Example output:
// "folder1 2023-01-01 15:00:00 user1", "folder2 this-is-folder-2 2023-01-01 15:00:10 user1"
func FoldersFormatted(user string, folders []Folder) []string {

	return nil
}
