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
		TheFiles    FilesManager
		Description string
	}

	// TheFiles interface defines the methods to manage files within a folder.
	FilesManager interface {
		AddFile(name, desc string) error
		RemoveFile(name string) error
		ListFiles(user, folder, by, order string) ([]string, error)
	}
)

// NewFolders creates a new Folders instance with initialized map.
func NewFolders() *Folders {

	return &Folders{
		List: make(map[string]*Folder),
	}
}

// AddFolder adds a new folder to the collection with the given name.
func (f *Folders) AddFolder(name, description string, filesManager FilesManager) error {

	if err := ValidateName(name); err != nil {

		return err
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
		TheFiles:    filesManager,
	}
	return nil
}

// FindFolder
func (f *Folders) FindFolder(name string) (*Folder, error) {
	folder, exists := f.List[name]
	if !exists {

		return nil, fmt.Errorf("Error: The '%s' doesn't exist", name)
	}

	return folder, nil
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

// ListFolders the folders list
// convert the sorting string value to the corresponding sorting type
func (f *Folders) ListFolders(user, theBy, theOrder string) ([]string, error) {

	var (
		by    sortBy    = 0
		order sortOrder = 0
	)

	if theBy == ByCreated {
		by = 1
	}

	if theOrder == DESC {
		order = 1
	}

	theFolders := f.GetSortedFolders(by, order)
	if len(theFolders) == 0 {

		return []string{}, fmt.Errorf("Warning: The '%s' doesn't have any folders.", user)
	}
	result := FoldersFormatted(user, theFolders)

	return result, nil
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
