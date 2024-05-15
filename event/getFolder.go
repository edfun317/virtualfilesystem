package event

import "iscoollab/filesystem/domain"

// getFolder retrieves a specific folder by name from a user's collection of folders.
// This function is used to access a folder's properties or to perform operations on the folder.
func getFolder(users *domain.Users, userName, folderName string) (*domain.Folder, error) {

	//retrieves the folder collection for the specified user from the Users object
	folders, err := users.GetUserFolders(userName)
	if err != nil {

		return nil, err
	}

	// get a folder from the folder collection
	folder, err := folders.FindFolder(folderName)
	if err != nil {
		return nil, err
	}

	return folder, nil
}
