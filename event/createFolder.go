package event

import (
	"iscoollab/filesystem/domain"
)

// CreateFolder add a new folder with a specified name and optional description
// to a user's collection of folders.
func CreateFolder(users *domain.Users, name, folder, desc string) error {

	//retrieves the folder collection for the specified user from the Users object
	folders, err := users.GetUserFolders(name)
	if err != nil {

		return err
	}

	//add a new folder with the given name and optional description.
	if err := folders.AddFolder(folder, desc); err != nil {

		return err
	}

	return nil
}
