package event

import "iscoollab/filesystem/domain"

// CreateFile add a new file with a specified name and optional description
func CreateFile(users *domain.Users, userName, folderName, fileName, desc string) error {

	//Find the specified folder
	folder, err := getFolder(users, userName, folderName)
	if err != nil {
		return err
	}

	//add a new file with the given name and optional description
	if err := folder.TheFiles.AddFile(fileName, desc); err != nil {
		return err
	}

	return nil
}
