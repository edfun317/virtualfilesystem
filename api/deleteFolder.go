package api

import (
	"fmt"
	"iscoollab/filesystem/domain"
	"iscoollab/filesystem/event"
	"os"
)

type (
	deleteFolder struct {
		Command
	}
)

// parseComands processes an input slice of strings to extract relevant command parameters.
func (d *deleteFolder) parseCommand(s []string) {
	length := len(s)

	if length >= 2 {
		d.User = s[1]
	}

	if length >= 3 {
		d.Folder = s[2]
	}
}

// Execute performs the delete folder operation
func (d *deleteFolder) execute(users *domain.Users) {

	if err := event.DeleteFolder(users, d.User, d.Folder); err != nil {

		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Printf("Delete '%s' successfully\n", d.User)

}
