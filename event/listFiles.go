package event

import (
	"errors"
	"iscoollab/filesystem/domain"
)

const (
	ByName    = "--sort-name"
	ByCreated = "--sort-created"
	ASC       = "asc"
	DESC      = "desc"
)

// ListFiles retrieves a sorted list of file names from a specified folder.
func ListFiles(users *domain.Users, userName, folderName, by, order string) ([]string, error) {

	err := errors.New("Usage: list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]")

	if userName == "" || folderName == "" {
		return nil, err
	}

	if by != "" || order != "" {

		if by == "" || order == "" {
			return nil, err
		}

		if by != ByName && by != ByCreated {
			return nil, err
		}

		if order != ASC && order != DESC {
			return nil, err
		}
	}

	//Find the specified folder
	folder, err := getFolder(users, userName, folderName)
	if err != nil {
		return nil, err
	}

	//list and sort the files based on the specified sort criteria.
	list, err := folder.TheFiles.ListFiles(userName, folderName, by, order)
	if err != nil {
		return nil, err
	}
	return list, nil
}
