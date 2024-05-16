package event

import (
	"errors"
	"iscoollab/filesystem/domain"
)

// DeleteFile removes a specified file from a folder's file collection.
func DeleteFile(users *domain.Users, userName, folderName, fileName string) error {

	if userName == "" || folderName == "" || fileName == "" {
		err := errors.New("Usage: delete-file [username] [foldername] [fileName]")
		return err
	}
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
