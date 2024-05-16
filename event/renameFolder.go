package event

import (
	"errors"
	"iscoollab/filesystem/domain"
)

// RenameFolder changes the name of an existing folder to a new name for a specified user.
func RenameFolder(users *domain.Users, userName, folderName, newFolder string) error {

	if userName == "" || folderName == "" || newFolder == "" {
		err := errors.New("Usage: rename-folder [username] [foldername] [new-folder-name]")
		return err
	}
	//retrieves the folder collection for the specified user from the Users object
	folders, err := users.GetUserFolders(userName)
	if err != nil {

		return err
	}

	//rename the specified folder. If the folder cannot be renamed
	// (e.g., if the folder does not exist or the new name is already in use), it returns an error.
	if err := folders.Rename(folderName, newFolder); err != nil {

		return err
	}

	return nil
}
