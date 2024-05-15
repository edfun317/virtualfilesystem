package event

import "iscoollab/filesystem/domain"

// DeleteFile removes a specified file from a folder's file collection.
func DeleteFile(users *domain.Users, userName, folderName, fileName string) error {

	//get a folder form folder collection
	folder, err := getFolder(users, userName, folderName)
	if err != nil {
		return err
	}

	//remove the specified file
	if err := folder.TheFiles.RemoveFile(fileName); err != nil {
		return err
	}

	return nil
}
