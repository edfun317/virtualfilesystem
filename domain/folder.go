package domain

import (
	"fmt"
	"sort"
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

	if err := ValidateName(name); err != nil {

		return fmt.Errorf("Error: The '%s' contain invalid chars", name)
	}
	//set mutex lock
	f.mux.Lock()
	//unlock
	defer f.mux.Unlock()

	if _, exists := f.List[name]; exists {

		return fmt.Errorf("Error: The '%s' has already existed.", name)
	}

	f.List[name] = &Folder{
		Name:        name,
		Created:     time.Now(),
		Description: description,
	}
	return nil
}

// RemoveFolder removes a folder from the collection by its name.
func (f *Folders) RemoveFolder(name string) error {

	f.mux.Lock()
	defer f.mux.Unlock()

	if _, exists := f.List[name]; !exists {
		return fmt.Errorf("Error: The '%s' doesn't exist", name)
	}

	delete(f.List, name)
	return nil
}

// Rename rename a folder
func (f *Folders) Rename(name, newName string) error {

	if err := ValidateName(name); err != nil {

		return fmt.Errorf("Error: The '%s' contain invalid chars", name)
	}

	f.mux.Lock()
	defer f.mux.Unlock()

	folder, exists := f.List[name]
	if !exists {

		return fmt.Errorf("Error: The '%s' doesn't exist", name)
	}

	if _, exists := f.List[newName]; exists {

		return fmt.Errorf("Error: The '%s' has already existed", newName)
	}
	folder.Name = newName
	f.List[newName] = folder

	delete(f.List, name)
	return nil
}

// GetSortedFolders returns a sorted slice of Folder based on the provided sort type and order.
func (f *Folders) GetSortedFolders(by sortBy, order sortOrder) []Folder {

	var folders []Folder
	for _, folder := range f.List {
		folders = append(folders, *folder)
	}

	// Sort the slice of folders according to the specified sort type and order.
	sort.Slice(folders, func(i, j int) bool {
		if by == sortByName {
			if order == ascending {
				return folders[i].Name < folders[j].Name
			}
			return folders[j].Name < folders[i].Name
		}
		if order == ascending {
			return folders[i].Created.Before(folders[j].Created)
		}
		return folders[j].Created.Before(folders[i].Created)
	})

	return folders
}

// FoldersFormatted formats a slice of Folder structures into a string slice
// for display purposes. The format includes the folder name, optional description,
// creation date, and associated user.
// Example output:
// "folder1 2023-01-01 15:00:00 user1", "folder2 this-is-folder-2 2023-01-01 15:00:10 user1"
func FoldersFormatted(user string, folders []Folder) []string {

	formatted := make([]string, 0, len(folders)) // Initialize with capacity to avoid resizing

	for _, folder := range folders {
		var folderStr string
		if folder.Description != "" {
			folderStr = fmt.Sprintf("%s %s %s %s", folder.Name, folder.Description, folder.Created.Format("2006-01-02 15:04:05"), user)
		} else {
			folderStr = fmt.Sprintf("%s %s %s", folder.Name, folder.Created.Format("2006-01-02 15:04:05"), user)
		}
		formatted = append(formatted, folderStr)
	}
	return formatted
}
