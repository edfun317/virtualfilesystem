package event

import (
	"fmt"
	"iscoollab/filesystem/domain"
	"os"
)

// CreateFolder add a new folder with a specified name and optional description
// to a user's collection of folders.
func CreateFolder(users *domain.Users, name, folder, desc string) {

	//retrieves the folder collection for the specified user from the Users object
	folders, err := users.GetUserFolders(name)
	if err != nil {

		//the error is printed to standard error
		fmt.Fprintln(os.Stderr, err)
		return
	}

	//add a new folder with the given name and description.
	if err := folders.AddFolder(folder, desc); err != nil {

		//the error is printed to standard error
		fmt.Fprintln(os.Stderr, err)
		return
	}

	//a success message is printed to standard output
	fmt.Printf("Create '%s' successfully.\n", folder)
}
