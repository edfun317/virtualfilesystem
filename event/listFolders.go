package event

import "iscoollab/filesystem/domain"

// ListFolders retrieves a sorted list of folder names from a user's collection of folders.
func ListFolders(users *domain.Users, userName, by, order string) ([]string, error) {

	//retrieves the folder collection for the specified user from the Users object
	user, err := users.GetUserFolders(userName)
	if err != nil {

		return nil, err
	}

	//list and sort the folders based on the specified sort criteria.
	list, err := user.ListFolders(userName, by, order)
	if err != nil {
		return nil, err
	}

	return list, nil
}
