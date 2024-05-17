package domain

import (
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"
)

const (
	ByName    = "--sort-name"
	ByCreated = "--sort-created"
	ASC       = "asc"
	DESC      = "desc"
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

	if err := ValidateName(name); err != nil {

		return fmt.Errorf("Error: The '%s' contain invalid chars", name)
	}

	//set mut lock
	f.mux.Lock()
	//unlock
	defer f.mux.Unlock()

	if _, exists := f.List[name]; exists {
		return fmt.Errorf("Error: The '%s' has already existed.", name)
	}

	f.List[name] = &File{
		Name:        name,
		Created:     time.Now(),
		Description: description,
	}

	return nil
}

// RemoveFile removes a file from the collection by its name
func (f *Files) RemoveFile(name string) error {

	f.mux.Lock()
	defer f.mux.Unlock()

	if _, exists := f.List[name]; !exists {
		return fmt.Errorf("Error: The '%s' doesn't exist", name)
	}
	delete(f.List, name)
	return nil
}

// ListFiles the files list.
// convert the sorting string value to the corresponding sorting type
func (f *Files) ListFiles(user, folder, theBy, theOrder string) ([]string, error) {

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

	theFiles := f.GetSortedFiles(by, order)
	if len(theFiles) == 0 {

		return []string{}, errors.New("Warning: The folder is empty")
	}

	result := FilesFormatted(user, folder, theFiles)

	return result, nil
}

// GetSortedFiles returns a sorted slice of File based
func (f *Files) GetSortedFiles(by sortBy, order sortOrder) []File {

	var files []File
	for _, file := range f.List {
		files = append(files, *file)
	}

	sort.Slice(files, func(i, j int) bool {
		if by == sortByName {
			if order == ascending {
				return files[i].Name < files[j].Name
			}
			return files[j].Name < files[i].Name
		}
		if order == ascending {
			return files[i].Created.Before(files[j].Created)
		}
		return files[j].Created.Before(files[i].Created)
	})
	return files
}

// FilesFormatted formats a slice of File structures into a string slice for display purposes.
// The format includes the file name, optional description, creation date, and associated user and folder.
// Example output: "file1 this-is-file1 2023-01-01 15:00:20 folder1 user1"
func FilesFormatted(user, folder string, files []File) []string {
	var formatted []string
	for _, file := range files {
		// Format the creation date and time as "YYYY-MM-DD HH:MM:SS".
		createdStr := file.Created.Format("2006-01-02 15:04:05")

		// Build the formatted string for each file. Include the description if it's not empty.
		var fileStr string
		if file.Description != "" {
			fileStr = fmt.Sprintf("%s %s %s %s %s", file.Name, file.Description, createdStr, folder, user)
		} else {
			fileStr = fmt.Sprintf("%s %s %s %s", file.Name, createdStr, folder, user)
		}
		formatted = append(formatted, fileStr)
	}
	return formatted
}
