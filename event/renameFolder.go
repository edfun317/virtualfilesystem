package event

import "iscoollab/filesystem/domain"

// RenameFolder changes the name of an existing folder to a new name for a specified user.
func RenameFolder(users *domain.Users, name, newName string) error {

	//retrieves the folder collection for the specified user from the Users object
	folders, err := users.GetUserFolders(name)
	if err != nil {

		return err
	}

	//rename the specified folder. If the folder cannot be renamed
	// (e.g., if the folder does not exist or the new name is already in use), it returns an error.
	if err := folders.Rename(name, newName); err != nil {

		return err
	}

	return nil
}
