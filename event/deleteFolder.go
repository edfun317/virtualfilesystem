package event

import (
	"errors"
	"iscoollab/filesystem/domain"
)

// DeleteFolder removes a specified folder from a user's folder collection.
func DeleteFolder(users *domain.Users, name, folder string) error {

	if name == "" || folder == "" {
		err := errors.New("Usage: delete-folder [username] [foldername]")
		return err
	}
	//retrieves the folder collection for the specified user from the Users object
	folders, err := users.GetUserFolders(name)
	if err != nil {

		return err
	}

	//remove the specified folder
	if err := folders.RemoveFolder(folder); err != nil {
		return err
	}

	return nil
}
