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
		Description string
	}
)
