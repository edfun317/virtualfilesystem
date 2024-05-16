package event

import (
	"errors"
	"iscoollab/filesystem/domain"
)

// CreateFolder add a new folder with a specified name and optional description
// to a user's collection of folders.
func CreateFolder(users *domain.Users, name, folder, desc string) error {

	if name == "" || folder == "" {
		err := errors.New("Usage: create-folder [username] [foldername] [description]?")
		return err
	}
	//retrieves the folder collection for the specified user from the Users object
	folders, err := users.GetUserFolders(name)
	if err != nil {

		return err
	}

	//add a new folder with the given name and optional description.
	if err := folders.AddFolder(folder, desc, domain.NewFiles()); err != nil {

		return err
	}

	return nil
}
